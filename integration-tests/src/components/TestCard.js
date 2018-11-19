import React from "react";
import PropTypes from "prop-types";
import Card from "@material-ui/core/Card";
import CardHeader from "@material-ui/core/CardHeader";
import CardActions from "@material-ui/core/CardActions";
import CardContent from "@material-ui/core/CardContent";
import Avatar from "@material-ui/core/Avatar";
import Grid from "@material-ui/core/Grid";
import IconButton from "@material-ui/core/IconButton";
import Typography from "@material-ui/core/Typography";
import PlayArrow from "@material-ui/icons/PlayArrow";
import Replay from "@material-ui/icons/Replay";
import Collapse from "@material-ui/core/Collapse";
import ExpandMoreIcon from "@material-ui/icons/ExpandMore";
import ExpandLessIcon from "@material-ui/icons/ExpandLess";
import { withStyles } from "@material-ui/core/styles";
import { styles } from "../styles/styles";
import classnames from "classnames";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText";
import PlayArrowWithSpinner from "./PlayArrowWithSpinner";
import Tooltip from "@material-ui/core/Tooltip";
import cardMediaWithIcon from "./CardMediaWithIcon";

let getStatusAsString = tests => {
    for (let i in tests) {
        if (tests[i].success === false) return "failure";
        if (tests[i].success === undefined) return "loading";
    }
    return "success";
};

class TestCard extends React.Component {
    state = { expanded: this.props.initiallyExpanded || false };

    handleExpandClick = () => {
        this.setState(state => ({ expanded: !state.expanded }));
    };

    render() {
        const { classes } = this.props;

        let status = getStatusAsString(this.props.tests);

        return (
            <Grid item md={3}>
                <Card className={classes.card}>
                    <CardHeader
                        avatar={
                            <Avatar
                                className={classes["statusAvatar" + status]}
                            >
                                {this.props.pathType}
                            </Avatar>
                        }
                        title={this.props.path}
                    />
                    {cardMediaWithIcon(classes, status)}
                    <CardContent className={classes.cardContent}>
                        <Typography
                            gutterBottom
                            variant="headline"
                            component="h2"
                        >
                            {status}
                        </Typography>
                        <Typography>{this.props.data.summary}</Typography>
                    </CardContent>
                    <CardActions>
                        <IconButton aria-label="Run tests">
                            <Tooltip title="Run All" placement="top">
                                <PlayArrow
                                    onClick={() =>
                                        this.props.tests.forEach(test =>
                                            this.props.runTest(
                                                this.props.pathType +
                                                    this.props.path,
                                                test.ID,
                                                test.test
                                            )
                                        )
                                    }
                                />
                            </Tooltip>
                        </IconButton>
                        <IconButton aria-label="Run failed tests">
                            <Tooltip title="Run failing" placement="top">
                                <Replay
                                    onClick={() =>
                                        this.props.tests.forEach(test => {
                                            if (!test.success)
                                                this.props.runTest(
                                                    this.props.pathType +
                                                        this.props.path,
                                                    test.ID,
                                                    test.test
                                                );
                                        })
                                    }
                                />
                            </Tooltip>
                        </IconButton>
                        <IconButton
                            className={classnames(classes.expand, {
                                [classes.expandOpen]: this.state.expanded
                            })}
                            onClick={this.handleExpandClick}
                            aria-expanded={this.state.expanded}
                            aria-label="Show more"
                        >
                            {this.state.expanded ? (
                                <ExpandLessIcon />
                            ) : (
                                <ExpandMoreIcon />
                            )}
                        </IconButton>
                    </CardActions>
                    <Collapse
                        in={this.state.expanded}
                        timeout="auto"
                        unmountOnExit
                    >
                        <CardContent>
                            <List component="nav">
                                {this.props.tests.map((test, i) => (
                                    <ListItem
                                        key={i}
                                        onClick={() => {
                                            this.props.runTest(
                                                this.props.pathType +
                                                    this.props.path,
                                                test.ID,
                                                test.test
                                            );
                                        }}
                                        button
                                    >
                                        <ListItemIcon>
                                            <PlayArrowWithSpinner
                                                classes={this.props.classes}
                                                success={test.success}
                                            />
                                        </ListItemIcon>
                                        <ListItemText
                                            inset
                                            primary={test.name}
                                        />
                                    </ListItem>
                                ))}
                            </List>
                        </CardContent>
                    </Collapse>
                </Card>
            </Grid>
        );
    }
}

TestCard.propTypes = {
    classes: PropTypes.object.isRequired,
    path: PropTypes.string.isRequired, // e.g. "/examples/services/hello"
    pathType: PropTypes.string.isRequired, // e.g. "get"
    data: PropTypes.object.isRequired, // e.g. {"summary":"HelloProxy says 'hello' in a form that is handled by the gateway proxy.","operationId":"HelloProxy","responses":{"200":{"description":"","schema":{"$ref":"#/definitions/protobufHelloResponse"}}},"parameters":[{"name":"hello_text","in":"query","required":false,"type":"string"}],"tags":["Exemplar"]}
    tests: PropTypes.array.isRequired, // [{name : accepts HelloProxy Request, success : true, test : () => {}}]
    runTest: PropTypes.func.isRequired, // callback for running the test
    initiallyExpanded: PropTypes.bool // open by default
};

export default withStyles(styles)(TestCard);
