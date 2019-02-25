const testEndpoint = require("../definitions/cliEndpoint");
const getpasswordsPositiveTest = require("../tests/getpasswordsPositiveTest");
var exec = require("child_process").exec;

describe("get/passwords", () => {
    test("PositiveTest", done => {
        let callback = (err, stdOut, stdError) => {
            try {
                expect(err).toBeFalsy();
                expect(() => JSON.parse(stdOut)).not.toThrow();
                if (getpasswordsPositiveTest.expectedOutput !== undefined) {
                    expect(JSON.parse(stdOut)).toEqual(
                        getpasswordsPositiveTest.expectedOutput
                    );
                }
                done();
            } catch (e) {
                done.fail(e + stdOut + err);
            }
        };

        let endpoint = testEndpoint.default + getpasswordsPositiveTest.path;
        let body = getpasswordsPositiveTest.requestBody || {};
        let command =
            "./src/cli/make_request.sh " +
            endpoint +
            " " +
            getpasswordsPositiveTest.method +
            " " +
            JSON.stringify(body);
        exec(command, callback);
    });
});
