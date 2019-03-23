/**
 * @jest-environment node
 */
const testEndpoint = require("../definitions/cliEndpoint");
const postchallengePositiveTest = require("../tests/postchallengePositiveTest");
var exec = require("child_process").exec;
const api = require("../api/api");

const MongoClient = require("mongodb").MongoClient;

let mockUser = {
    user: {
        first: "David",
        last: "Goldstein",
        email: "david@david.com"
    },
    auth: {
        dn: "postchallengetest@integration-tests.com",
        authpassword: "poiw83vliuaywelfiauwheflkajsdhdrvl938lwu3iohf",
        accesstoken: "",
        failedlogins: 0,
        authquestions: [
            {
                q: "what is your pet's name?",
                a: "penny"
            }
        ],
        knownips: ["172.42.64.6"]
    },
    logins: [
        {
            timestamp: 1541867160,
            location: {
                "ip": "172.42.64.6",
                "type": "ipv4",
                "countrycode": "US",
                "countryname": "United States",
                "regioncode": "VA",
                "regionname": "Virginia",
                "city": "Charlottesville",
                "zip": "22901",
                "latitude": 38.0936,
                "longitude": -78.5611
            }
        }
    ],
    passwords: "ewogIC"
};

let lockedOutUser = {
    user: {
        first: "Locked",
        last: "Out",
        email: "locked@out.com"
    },
    auth: {
        dn: "postchallengetestlockedout@integration-tests.com",
        authpassword: "poiw83vliuaywelfiauwheflkajsdhdrvl938lwu3iohf",
        accesstoken: "",
        failedlogins: 10,
        authquestions: [
            {
                q: "what is your pet's name?",
                a: "penny"
            }
        ],
        knownips: ["172.42.64.6"]
    },
    logins : [],
    passwords : "sdfsdf",
};

let collection, client;
describe("post/challenge", () => {
    beforeEach(done => {
        // conect to mongodb
        const url = "mongodb://localhost:27017";
        client = new MongoClient(url);
        client.connect(function(err) {
            if (err) {
                console.error(err);
                done();
            }
            const db = client.db("passwords");
            collection = db.collection("passwords");
            // insert mock data into mongo
            collection.insertMany([mockUser, lockedOutUser], err => {
                if (err) console.error(err);
                done();
            });
        });
    });

    afterEach(done => {
        collection.deleteMany(
            {},
            undefined,
            err => {
                if (err) console.error(err);
                client.close();
                done();
            }
        );
    });

    let endpoint =
        testEndpoint.default +
        postchallengePositiveTest.path +
        "?user=" +
        mockUser.auth.dn;

    test("retrieve new challenge", done => {
        let body = {
            location: {
                ...mockUser.logins[0].location,
                countryCode : mockUser.logins[0].location.countrycode,
            }
        };

        api.post(endpoint, body).then(({ error, success, data }) => {
            try {
                if (error) throw new Error(error.response.data.error);
                expect(success).toBe(true);
                expect(error).toBeFalsy();
                if (postchallengePositiveTest.expectedOutput !== undefined) {
                    expect(data).not.toBeUndefined()
                    expect(data.user).toEqual(mockUser.user)
                    
                }
                done();
            } catch (e) {
                console.log(error, success, data)
                done.fail(e);
            }
        });
    });
    test("generates user defined challenge if location is not known", done => {
        let newIp = "172.0.0.3";
        // give new ip address
        let body = {
            location : {
              ...mockUser.logins[0].location,
              countryCode : mockUser.logins[0].location.countrycode,
              ip : newIp,                
            }
        }

        api.post(endpoint, body).then(({ error, success, data }) => {
            try {
                if (error) throw new Error(error.response.data.error);
                expect(success).toBe(true);
                expect(error).toBeFalsy();
                if (postchallengePositiveTest.expectedOutput !== undefined) {
                    expect(data).not.toBeUndefined()
                    expect(data.error).toEqual('Login Unsuccessful')
                    expect(data.userQuestion).toEqual(mockUser.auth.authquestions[0].q)
                    
                }
                done();
            } catch (e) {
                console.log(error, success, data)
                done.fail(e);
            }
        });
    })
    test("retrieve new challenge if location is now know but know user-defined auth", done => {        
        let newIp = "172.0.0.3";
        // give new ip address, but this time with user defined answer
        let body = {
            "userQuestionResponse" : {
                q : mockUser.auth.authquestions[0].q,
                a : mockUser.auth.authquestions[0].a,
            },
            location : {
              ...mockUser.logins[0].location,
              countryCode : mockUser.logins[0].location.countrycode,
              ip : newIp,                
            }
        }

        api.post(endpoint, body).then(({ error, success, data }) => {
            try {
                if (error) throw new Error(error.response.data.error);
                expect(success).toBe(true);
                expect(error).toBeFalsy();
                if (postchallengePositiveTest.expectedOutput !== undefined) {
                    expect(data).not.toBeUndefined()
                    expect(data.user).toEqual(mockUser.user)
                    
                }
                done();
            } catch (e) {
                console.log(error, success, data)
                done.fail(e);
            }
        });
        
    })
    test("gets locked out if more than 5 incorrect logins", done => {
        let body = {
            location : {
              ...mockUser.logins[0].location,
              countryCode : mockUser.logins[0].location.countrycode,
              ip : "127.4.7.4",
            }
        }

        let lockedOutUserEndpoint = 
        testEndpoint.default +
        postchallengePositiveTest.path +
        "?user=" +
        lockedOutUser.auth.dn
        
        // make request to test that we're locked out
        api.post(lockedOutUserEndpoint, body).then(({error, success, data}) => {
            expect(error.response.data.error).toEqual(`'${lockedOutUser.auth.dn}' is locked out. Please contact an administrator to regain access.`)
            done()
        })               
    })
});
