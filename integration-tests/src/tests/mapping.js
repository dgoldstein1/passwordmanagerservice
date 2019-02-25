// mapping.js
import { postchallengePositiveTest } from "./postchallengePositiveTest";
import { getpasswordsPositiveTest } from "./getpasswordsPositiveTest";
import { postpasswordsPositiveTest } from "./postpasswordsPositiveTest";

export default {
    "post/challenge": {
        postchallengePositiveTest: {
            name: "PositiveTest",
            ID: "postchallengePositiveTest",
            test: postchallengePositiveTest
        }
    },
    "get/passwords": {
        getpasswordsPositiveTest: {
            name: "PositiveTest",
            ID: "getpasswordsPositiveTest",
            test: getpasswordsPositiveTest
        }
    },
    "post/passwords": {
        postpasswordsPositiveTest: {
            name: "PositiveTest",
            ID: "postpasswordsPositiveTest",
            test: postpasswordsPositiveTest
        }
    }
};
