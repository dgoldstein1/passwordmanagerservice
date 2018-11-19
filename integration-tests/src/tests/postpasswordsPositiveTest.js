// replace code here
import api from "../api/api";
import endpoint from "../definitions/endpoint";
import _ from "lodash";
let params = {};
let path = "/passwords" + api.paramsToUri(params);
let requestBody = {};
let method = "POST";
let expectedOutput =undefined;
// method run during testing
let postpasswordsPositiveTest = function() {
    return api[method.toLowerCase()](endpoint + path, requestBody);
};
// footer, configured this way for testing
export { postpasswordsPositiveTest, method, requestBody, expectedOutput, path };
