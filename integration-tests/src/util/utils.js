// utils.js

/**
 * @param {swagger} swagger json file
 * @return {array of json in form}

   classes: PropTypes.object.isRequired,
   path : PropTypes.string.isRequired, // e.g. "/examples/services/hello"
   pathType : PropTypes.string.isRequired, // e.g. "get"
   data : PropTypes.object.isRequired
 * 
 **/
let swaggerToTestCardArray = swagger => {
    let toReturn = [];
    for (let path in swagger.paths) {
        for (let pathType in swagger.paths[path]) {
            toReturn.push({
                path,
                pathType,
                data: swagger.paths[path][pathType]
            });
        }
    }
    return toReturn;
};

/**
 * gets header color from the state of all tests
 * if any tests are undefined, is yellow
 * then, if any tests are red, is red
 * else is green
 *
 * @param {json} classes
 * @param {json} test object
 * @return {string} className
 **/
let getHeaderClass = (classes, tests) => {
    let someTestsHaveFailed = false;
    // loop through tests
    for (let endpt in tests) {
        for (let test in tests[endpt]) {
            let t = tests[endpt][test];
            // if any test is undefined, return loading
            if (t.success === undefined) return classes.appBarLoading;
            if (t.success === false) someTestsHaveFailed = true;
        }
    }
    if (someTestsHaveFailed) return classes.appBarFailure;
    return classes.appBarSuccess;
};

/**
 * @param {json} {key : {data : 1}, key2 : {data : 2}}
 * return {array} [{data : 1}, {data : 2}]
 **/
let JSONtoArray = json => {
    let toReturn = [];
    for (let i in json) {
        toReturn.push(json[i]);
    }
    return toReturn;
};

export { swaggerToTestCardArray, JSONtoArray, getHeaderClass };
