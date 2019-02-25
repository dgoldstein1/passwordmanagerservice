const testEndpoint = require("../definitions/cliEndpoint");
const postchallengePositiveTest = require("../tests/postchallengePositiveTest");
var exec = require("child_process").exec;

describe("post/challenge", () => {
    test("PositiveTest", done => {
        let callback = (err, stdOut, stdError) => {
            try {
                expect(err).toBeFalsy();
                expect(() => JSON.parse(stdOut)).not.toThrow();
                if (postchallengePositiveTest.expectedOutput !== undefined) {
                    expect(JSON.parse(stdOut)).toEqual(
                        postchallengePositiveTest.expectedOutput
                    );
                }
                done();
            } catch (e) {
                done.fail(e + stdOut + err);
            }
        };

        let endpoint = testEndpoint.default + postchallengePositiveTest.path;
        let body = postchallengePositiveTest.requestBody || {};
        let command =
            "./src/cli/make_request.sh " +
            endpoint +
            " " +
            postchallengePositiveTest.method +
            " " +
            JSON.stringify(body);
        exec(command, callback);
    });
});
