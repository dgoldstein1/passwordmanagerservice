// replace code here
import api from "../api/api";
import endpoint from "../definitions/endpoint";
import _ from "lodash";
let params = {};
let path = "/challenge" + api.paramsToUri(params);
let requestBody = {
    userQuestionResponse: { q: "Dqauxvt", a: "FpCNpz" },
    location: {
        ip: "j",
        type: "qKwfkGYBHZU",
        countryCode: "kCRVLXeru",
        countryName: "esQtMkDZ",
        regionCode: "Cg",
        regionName: "ZtvpmoCbo",
        city: "CcuSIh",
        zip: "doXpTswpZgW",
        latitude: -829108894302.208,
        longitude: 423995532482.9696
    }
};
let method = "POST";
let expectedOutput = {
    error: "UpvvqRevKvN",
    userQuestion: "fUIKzKwQd",
    logins: [
        {
            timestamp: "ImqwbUEwdZH",
            location: {
                ip: "QKwBhMokMR",
                type: "dqvypmPK",
                countryCode: "LsErbzN",
                countryName: "PtmHIgLf",
                regionCode: "F",
                regionName: "OzMOBF",
                city: "HQxlX",
                zip: "BFuLIZfo",
                latitude: -481622473519.9232,
                longitude: -451971742747.8528
            }
        }
    ],
    user: { first: "dtO", last: "gkia", email: "XzRHUWwlCH" },
    challenge: "gUoXAEeN"
};
// method run during testing
let postchallengePositiveTest = function() {
    return api[method.toLowerCase()](endpoint + path, requestBody);
};
// footer, configured this way for testing
export { postchallengePositiveTest, method, requestBody, expectedOutput, path };
