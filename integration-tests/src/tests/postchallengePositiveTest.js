// replace code here
import api from "../api/api";
import endpoint from "../definitions/endpoint";
import _ from "lodash";
let params = {};
let path = "/challenge" + api.paramsToUri(params);
let requestBody = {};
let method = "POST";
let expectedOutput = {"code": 2, "error": "not implemented", "message": "not implemented"};
// method run during testing
let postchallengePositiveTest = function() {
    return api[method.toLowerCase()](endpoint + path, requestBody);
};
// footer, configured this way for testing
export { postchallengePositiveTest, method, requestBody, expectedOutput, path };
