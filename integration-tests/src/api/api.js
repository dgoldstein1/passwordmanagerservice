// api.js

// fetching library
const axios = require("axios");
require("whatwg-fetch");

/**
 * an axios helper to make requests
 **/

let headers = { "Content-Type": "application/json" };

let get = function(url) {
    return axios
        .get(url, headers)
        .then(res => ({ url: url, success: true, data: res.data }))
        .catch(error => ({ url: url, success: false, error: error }));
};

let patch = function(url, body = {}) {
    return axios
        .patch(url, body, headers)
        .then(res => ({ url: url, success: true, data: res.data }))
        .catch(error => ({ url: url, success: false, error: error }));
};

let post = function(url, body = {}) {
    return axios
        .post(url, body, headers)
        .then(res => ({ url: url, success: true, data: res.data }))
        .catch(error => ({ url: url, success: false, error: error }));
};

let del = function(url, body = {}) {
    return axios
        .delete(url, body, headers)
        .then(res => ({ url: url, success: true, data: res.data }))
        .catch(error => ({ url: url, success: false, error: error }));
};

/////////////////
// uri helpers //
/////////////////

/**
 * takes in params to add them to url
 * @param {json} params (i.e. {pagenumber : 1})
 * @return {string} params as url query (i.e. ?pagenumber=1)
 **/
const paramsToUri = function(params) {
    let uri = "";

    for (let param in params) {
        uri += `&${param}=${params[param]}`;
    }
    // if we've added things, we need to put a question mark on the beginning
    if (uri.length > 0) uri = "?" + uri.substr(1, uri.length);

    return uri;
};

module.exports = {
    get,
    patch,
    post,
    del,
    paramsToUri
};
