import React from "react";
import PropTypes from "prop-types";
import classNames from "classnames";
import AppBar from "@material-ui/core/AppBar";
import Button from "@material-ui/core/Button";
import CssBaseline from "@material-ui/core/CssBaseline";
import Grid from "@material-ui/core/Grid";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import { withStyles } from "@material-ui/core/styles";
import { styles, theme } from "../styles/styles";
import { MuiThemeProvider } from "@material-ui/core/styles";
import { withTheme } from "@material-ui/core/styles";
import TestCard from "./TestCard";

import mapping from "../tests/mapping";

import {
    JSONtoArray,
    swaggerToTestCardArray,
    getHeaderClass
} from "../util/utils";

import swagger from "../definitions/swagger";

class App extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            tests: this.props.testMapping || mapping, // show test mapping if passed
            loading: false
        };
        this._runTest = this._runTest.bind(this);
        this._runAllTests = this._runAllTests.bind(this);
    }

    /**
     * runs all tests in the mapping file
     * @param {bool} run only failing tests?
     **/
    _runAllTests = (failingOnly = false) => {
        for (let path in this.state.tests) {
            for (let test in this.state.tests[path]) {
                // declare temp json test object
                let t = this.state.tests[path][test];
                // do not run if failing only and is success
                if (failingOnly && t.success === true) continue;
                // else run test
                this._runTest(path, t.ID, t.test);
            }
        }
    };

    /**
     * actions done before the app is loaded
     * - set app name on tab to swagger definition name
     * - run all tests
     **/
    componentDidMount() {
        document.title = swagger.info.title;
        // run all tests
        this._runAllTests();
    }

    /**
     * callback for running a specific test
     * @param {string} testPath. Path to the card (e.g. "get/examples/services/hello")
     * @param {string} testId specific id within the card "randomId1"
     * @param {function} test which returns Promise{success : true}
     * sets state of test using set state
     **/
    _runTest = (testPath, testName, testFunction) => {
        // set test as loading
        this.setState(currState => {
            currState.tests[`${testPath}`][`${testName}`][
                "success"
            ] = undefined;
            return { currState };
        });
        // run test
        testFunction().then(res => {
            // set updated state of test
            this.setState(currState => {
                currState.tests[`${testPath}`][`${testName}`]["success"] =
                    res.success;
                return { currState };
            });
        });
    };

    render() {
        const { classes } = this.props;

        return (
            <MuiThemeProvider theme={theme}>
                <CssBaseline />
                <AppBar
                    position="static"
                    className={getHeaderClass(classes, this.state.tests)}
                >
                    <Toolbar>
                        <Typography
                            variant="title"
                            color="inherit"
                            align="center"
                            noWrap
                        >
                            {`${swagger.info.title} ${
                                swagger.info.version
                            } Integration Test Suite`}
                        </Typography>
                    </Toolbar>
                </AppBar>
                <main>
                    {/* Hero unit */}
                    <div className={classes.heroUnit}>
                        <div className={classes.heroContent}>
                            <Typography
                                variant="display3"
                                align="center"
                                color="textPrimary"
                                gutterBottom
                            >
                                Service Integration Tests
                            </Typography>
                            <Typography
                                variant="title"
                                align="center"
                                color="textSecondary"
                                paragraph
                            >
                                Tests automatically generated from a swagger
                                definition file.
                            </Typography>
                            <div className={classes.heroButtons}>
                                <Grid container spacing={16} justify="center">
                                    <Grid item>
                                        <Button
                                            variant="outlined"
                                            onClick={() => this._runAllTests()}
                                        >
                                            Run All Tests
                                        </Button>
                                    </Grid>
                                    <Grid item>
                                        <Button
                                            variant="outlined"
                                            onClick={() =>
                                                this._runAllTests(true)
                                            }
                                        >
                                            Run Failing Tests
                                        </Button>
                                    </Grid>
                                </Grid>
                            </div>
                        </div>
                    </div>
                    <div
                        className={classNames(classes.layout, classes.cardGrid)}
                    >
                        {/* End hero unit */}
                        <Grid container justify="center" spacing={40}>
                            {swaggerToTestCardArray(swagger).map((card, i) => (
                                <TestCard
                                    key={i}
                                    runTest={this._runTest}
                                    tests={JSONtoArray(
                                        this.state.tests[
                                            `${card.pathType}${card.path}`
                                        ]
                                    )}
                                    {...card}
                                />
                            ))}
                        </Grid>
                    </div>
                </main>
                {/* Footer */}
                <footer className={classes.footer}>
                    <Typography variant="title" align="center" gutterBottom>
                        Copyright © 2018{" "}
                        <a
                            style={{ textDecoration: "none", color: "black" }}
                            href="http://deciphernow.com/"
                        >
                            {" "}
                            Decipher Technology Studios{" "}
                        </a>
                    </Typography>
                    <Typography
                        variant="subheaidng"
                        align="center"
                        color="textSecondary"
                        component="p"
                    >
                        <a
                            style={{ textDecoration: "none", color: "black" }}
                            href="https://github.com/dgoldstein1/integration-test-generator"
                        >
                            https://github.com/dgoldstein1/integration-test-generator
                        </a>
                    </Typography>
                </footer>
                {/* End footer */}
            </MuiThemeProvider>
        );
    }
}

App.propTypes = {
    classes: PropTypes.object.isRequired
};

export default withTheme(theme)(withStyles(styles)(App));
