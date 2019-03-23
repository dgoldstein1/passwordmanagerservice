// replace code here
import api from "../api/api";
import endpoint from "../definitions/endpoint";
import _ from "lodash";
let params = {};
let path = "/passwords" + api.paramsToUri(params);
let requestBody = {};
let method = "GET";
let expectedOutput = { passwords: "IF" };
// method run during testing
let getpasswordsPositiveTest = function() {
    return api[method.toLowerCase()](endpoint + path, requestBody);
};
// footer, configured this way for testing
export { getpasswordsPositiveTest, method, requestBody, expectedOutput, path };
