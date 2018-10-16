#!/usr/bin/python
#
# The init script for the passwordservice service. Save to /etc/init.d
#
#
# chkconfig: - 20 70
# description: passwordservice-X.X server process
#
#
### BEGIN INIT INFO
# Provides: passwordservice-X.X
# Description: passwordservice-X.X server process
### END INIT INFO

import sys, os, subprocess, re, time, signal, logging
from pwd import getpwnam
from grp import getgrnam

logging.basicConfig(format='%(levelname)s:%(message)s', level=logging.INFO)

USER = "passwordservice"
GROUP = "services"

VERSION_MAJOR = "1"
VERSION_MINOR = "1"
VERSION_PATCH = "0"

SERVICE = "passwordservice-{0}.{1}".format(VERSION_MAJOR, VERSION_MINOR)
INSTALL_PATH = os.path.join('opt','services',SERVICE)
EXE = os.path.join(INSTALL_PATH,'bin','passwordservice-bin')
CONFIG = os.path.join(INSTALL_PATH,'etc','settings.toml')
LOCKFILE = os.path.join(INSTALL_PATH,'lockfile')
PIDFILE = os.path.join(INSTALL_PATH,'pidfile')
LOG_OUT = "/var/log/{0}.log".format(SERVICE)
LOG_ERR = "/var/log/{0}.err".format(SERVICE)


def change_owner(path, uid, gid, perms):
    """
    Set permissions on file or dir. uid and perms must be int type, eg. 0o755.
    Group is ignored with -1.
    """
    logging.debug('In change_owner for %s', path)
    statinfo = os.stat(path)
    puid = statinfo.st_uid
    pgid = statinfo.st_gid
    pmode = statinfo.st_mode
    if uid != puid or (gid != -1 and gid != pgid):
        os.chown(path, uid, gid)
    if perms != oct(pmode & 0o777):
        os.chmod(path, perms)


def create_dir_not_exists(path):
    logging.debug('In create_dir_not_exists for %s', path)
    if not os.path.exists(path):
        logging.info('Creating direcotry: %s', path)
        os.makedirs(path)


def lock():
    logging.debug('In lock')
    if not locked():
        open(LOCKFILE, 'w').close()
    else:
        pid = get_my_pid()
        if pid < 0:
            # process not running and no pidfile
            return 0
        else:
            if is_process_running(pid):
                logging.warning('Service %s is already running', SERVICE)
                sys.exit(1)
            else:
                if os.path.exists(PIDFILE):
                    logging.warning('Removing stale pidfile')
                    removepidfile()            


def locked():
    logging.debug('In locked')
    if os.path.exists(LOCKFILE):
        logging.debug('- lockfile found')
        return True
    logging.debug('- no lockfile')
    return False


def touch_file(path):
    logging.debug('In touch_file for %s', path)
    if not os.path.exists(path):
        f = open(path, "w")
        f.close()


def unlock():
    logging.debug('In unlock')
    if os.path.exists(LOCKFILE):
        os.remove(LOCKFILE)


def start():
    logging.debug('In start')

    removepidfile()

    # Get target uid/gid.
    uid = getpwnam(USER).pw_uid
    gid = getgrnam(GROUP).gr_gid

    # Capture root's environment in dictionary. Pass to Popen.
    env_for_proc = dict()
    for k in os.environ.keys():
        env_for_proc[k] = os.getenv(k, "")

    # Change permissions on configuration file.
    change_owner(CONFIG, uid, gid, 0o640)

    # Change ownership of logging directory and file, if the file exists.
    logging_dir = os.path.split(LOG_OUT)[0]
    create_dir_not_exists(logging_dir)
    change_owner(logging_dir, uid, gid, 0o750)
    if os.path.exists(LOG_OUT):
        change_owner(LOG_OUT, uid, gid, 0o640)

    # Change ownership of pidfile.
    touch_file(PIDFILE)
    change_owner(PIDFILE, uid, gid, 0o640)

    # Change group/user (note order).
    os.setgid(gid)
    os.setuid(uid)

    logging.info('Starting %s', SERVICE)
    f = open(LOG_OUT, "a")

    command_client = [EXE, '--config', CONFIG]
    try:
        ps = subprocess.Popen(command_client, stdout=f, stderr=f, env=env_for_proc)
        with open(PIDFILE, "w") as pf:
            pf.write(str(ps.pid) + '\n')
    except Exception as e:
        print e
        removepidfile()
        unlock()
        sys.exit(1)


def stop():
    logging.debug('In stop')
    logging.info('Stopping %s', SERVICE)
    try:
        pid = get_my_pid()
        if pid < 0:
            # process not running and no pidfile
            return 0
        else:
            # either pidfile or a process
            if not is_process_running(pid):
                # no running process, cleanup
                if os.path.exists(PIDFILE):
                    logging.warning('Removing stale pidfile')
                    removepidfile()
                return 0
            else:
                # there is a running process, kill it
                os.kill(int(pid), signal.SIGTERM)
                logging.info('SIGTERM sent to process.')
                time.sleep(5)
                removepidfile()
                return 0
    except Exception as e:
        print e
        logging.error('Could not kill process')
        sys.exit(1)
        return 1


def restart():
    logging.debug('In restart')
    r = stop()
    if r == 0:
        lock()
        start()
    else:
        sys.exit(1)


def status():
    logging.debug('In status')
    if not locked():
        logging.info('%s is not locked or running', SERVICE)
        return 3
    else:
        if not os.path.exists(PIDFILE):
            logging.info('%s is not running', SERVICE)
            return 3
        pid = get_pid_from_pidfile(PIDFILE)
        if not is_process_running(pid):
            logging.warning('A pidfile exists but %s is not running', SERVICE)
            logging.warning('Removing stale pidfile for %s', SERVICE)
            removepidfile()
            return 1
        logging.info('%s is running (pid %s)', SERVICE, pid)
        return 0


def removepidfile():
    logging.debug('In removepidfile')
    if os.path.exists(PIDFILE):
        os.remove(PIDFILE)


def get_pid_from_pidfile(path):
    logging.debug('In get_pid_from_pidfile for %s', path)
    with open(path, 'r') as f:
        pid = f.readline().split('\n')[0]
        if len(pid) > 0:
            return pid
        else:
            return -1


def get_pid_for_process():
    logging.debug('In get_pid_from_process')
    child = subprocess.Popen(['pgrep',SERVICE], stdout=subprocess.PIPE, shell=False)
    results = child.communicate()
    if len(results) > 0:
        return results[0]
    else:
        return -1


def get_my_pid():
    logging.debug('In get_my_pid')
    if not locked():
        return get_pid_for_process()
    else:
        if not os.path.exists(PIDFILE):
            return -1
        else:
            return get_pid_from_pidfile(PIDFILE)


def is_process_running(process_id):
    logging.debug('In is_process_running')
    if len(process_id) == 0:
        return False
    try:
        os.kill(int(process_id), 0)
        return True
    except OSError:
        return False


def is_service_running():
    if not os.path.exists(PIDFILE):
        return False
    else:
        pid = get_pid_from_pidfile(PIDFILE)
    return is_process_running(pid)


def is_absolute_path(path):
    logging.debug('In is_absolute_path for %s', path)
    return os.path.isabs(path)


def shell_source(script):
    source_file = open(script, "r")
    source_dict = {}
    prefix = "export "
    for line in source_file:
        if line.lower().startswith(prefix):
            line = line[len(prefix):]
            line_parts = line.split("=", 1)
            line_key = line_parts[0]
            line_value = line_parts[1].rstrip("\r\n ")
            line_value = line_value.lstrip("\"")
            line_value = line_value.rstrip("\"")
            line_value = line_value.lstrip("\'")
            line_value = line_value.rstrip("\'")
            line_value = os.path.expandvars(line_value)
            source_dict[line_key] = line_value
            os.environ.update(source_dict)
    source_file.close()


# Script entry point.
if __name__ == '__main__':
    try:
        if len(sys.argv) == 1:
            raise ValueError
        create_dir_not_exists(os.path.split(LOG_OUT)[0])
        command = str(sys.argv[1]).strip().lower()
        if command == 'start':
            if not is_service_running():
                lock()
                start()
            sys.exit(0)
        elif command == 'stop':
            stop()
            unlock()
            sys.exit(0)
        elif command == 'restart' or command == 'force-reload':
            restart()
            sys.exit(0)
        elif command == 'status':
            ok = status()
            sys.exit(ok)
        else:
            raise ValueError
    except (SystemExit):
        pass
    except (ValueError):
        print >> sys.stderr, "Usage: {0} [start|stop|restart|status]".format(SERVICE)
        sys.exit(2)
    except:
        # Other errors
        extype, value = sys.exc_info()[:2]
        print >> sys.stderr, "ERROR: %s (%s)" % (extype, value)
        sys.exit(1)        
