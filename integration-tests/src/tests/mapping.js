// mapping.js
import { postpasswordsPositiveTest } from "./postpasswordsPositiveTest";

export default {
    "post/passwords": {
        postpasswordsPositiveTest: {
            name: "PositiveTest",
            ID: "postpasswordsPositiveTest",
            test: postpasswordsPositiveTest
        }
    }
};
