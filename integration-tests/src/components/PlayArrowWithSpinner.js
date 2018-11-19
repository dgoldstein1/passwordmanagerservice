// PlayArrowWithSpinner.js

import React from "react";
import PropTypes from "prop-types";
import { withStyles } from "@material-ui/core/styles";
import PlayArrow from "@material-ui/icons/PlayArrow";
import CheckIcon from "@material-ui/icons/Check";
import ErrorIcon from "@material-ui/icons/Error";
import Button from "@material-ui/core/Button";
import CircularProgress from "@material-ui/core/CircularProgress";

import { styles } from "../styles/styles";

/**
 * Play arrow material-ui icon with loading spinner
 **/

// injected styles

class PlayArrowWithSpinner extends React.Component {
    render() {
        let { success, classes } = this.props;

        /**
         * helper for getting classes
         * @param {bool} success
         * @param {json} class names
         * @return {json} {button, backgroundColor} classes
         **/
        let _getButtonClass = function(success, classes) {
            if (success === true) return classes.buttonSuccess;
            if (success === false) return classes.buttonFailure;
            return classes.buttonLoading;
        };

        return (
            <div className={classes.root}>
                <div className={classes.wrapper}>
                    <Button
                        variant="fab"
                        className={_getButtonClass(success, classes)}
                    >
                        {success === true && <CheckIcon />}
                        {success === undefined && <PlayArrow />}
                        {success === false && <ErrorIcon />}
                    </Button>
                    {success === undefined && (
                        <CircularProgress
                            size={68}
                            className={classes.fabProgress}
                        />
                    )}
                </div>
            </div>
        );
    }
}

PlayArrowWithSpinner.propTypes = {
    classes: PropTypes.object.isRequired, // classes for styles
    success: PropTypes.bool // if undefined, means it's currently loading
};

export default withStyles(styles)(PlayArrowWithSpinner);
