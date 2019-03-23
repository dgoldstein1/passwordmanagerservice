const testEndpoint = require("../definitions/cliEndpoint");
const postpasswordsPositiveTest = require("../tests/postpasswordsPositiveTest");
var exec = require("child_process").exec;

describe.skip("post/passwords", () => {
    test("PositiveTest", done => {
        let callback = (err, stdOut, stdError) => {
            try {
                expect(err).toBeFalsy();
                expect(() => JSON.parse(stdOut)).not.toThrow();
                if (postpasswordsPositiveTest.expectedOutput !== undefined) {
                    expect(JSON.parse(stdOut)).toEqual(
                        postpasswordsPositiveTest.expectedOutput
                    );
                }
                done();
            } catch (e) {
                done.fail(e + stdOut + err);
            }
        };

        let endpoint = testEndpoint.default + postpasswordsPositiveTest.path;
        let body = postpasswordsPositiveTest.requestBody || {};
        let command =
            "./src/cli/make_request.sh " +
            endpoint +
            " " +
            postpasswordsPositiveTest.method +
            " " +
            JSON.stringify(body);
        exec(command, callback);
    });
});
