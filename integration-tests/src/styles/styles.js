// styles.js

import { createMuiTheme } from "@material-ui/core/styles";

// colors
import red from "@material-ui/core/colors/red";
import green from "@material-ui/core/colors/green";
import yellow from "@material-ui/core/colors/yellow";

const LOADING_COLOR = yellow[800];
const SUCCESS_COLOR = green[500];
const FAILURE_COLOR = red[500];

// decipher theme
const theme = createMuiTheme({
    palette: {
        primary: green,
        secondary: green
    }
});

const styles = theme => ({
    appBarSuccess: {
        position: "relative",
        backgroundColor: SUCCESS_COLOR
    },
    appBarFailure: {
        position: "relative",
        backgroundColor: FAILURE_COLOR
    },
    appBarLoading: {
        position: "relative",
        backgroundColor: LOADING_COLOR
    },
    icon: {
        marginRight: theme.spacing.unit * 2
    },
    heroUnit: {
        backgroundColor: theme.palette.background.paper
    },
    heroContent: {
        maxWidth: 600,
        margin: "0 auto",
        padding: `${theme.spacing.unit * 8}px 0 ${theme.spacing.unit * 6}px`
    },
    heroButtons: {
        marginTop: theme.spacing.unit * 4
    },
    layout: {
        width: "auto",
        marginLeft: theme.spacing.unit * 3,
        marginRight: theme.spacing.unit * 3,
        [theme.breakpoints.up(1100 + theme.spacing.unit * 3 * 2)]: {
            width: 1100,
            marginLeft: "auto",
            marginRight: "auto"
        }
    },
    cardGrid: {
        padding: `${theme.spacing.unit * 8}px 0`
    },
    card: {
        height: "100%",
        display: "flex",
        flexDirection: "column"
    },
    cardMedia: {},
    cardContent: {
        flexGrow: 1
    },
    footer: {
        backgroundColor: theme.palette.background.paper,
        padding: theme.spacing.unit * 6
    },
    statusAvatarloading: {
        // loading
        backgroundColor: LOADING_COLOR,
        height: "60px",
        width: "60px"
    },
    statusAvatarfailure: {
        // test failed
        backgroundColor: FAILURE_COLOR,
        height: "60px",
        width: "60px"
    },
    statusAvatarsuccess: {
        // success!
        backgroundColor: SUCCESS_COLOR,
        height: "60px",
        width: "60px"
    },
    progress: {
        margin: theme.spacing.unit * 2
    },
    root: {
        display: "flex",
        alignItems: "center"
    },
    wrapper: {
        margin: theme.spacing.unit,
        position: "relative"
    },
    buttonSuccess: {
        backgroundColor: SUCCESS_COLOR,
        "&:hover": {
            backgroundColor: SUCCESS_COLOR
        }
    },
    buttonLoading: {
        backgroundColor: LOADING_COLOR,
        "&:hover": {
            backgroundColor: LOADING_COLOR
        }
    },
    buttonFailure: {
        backgroundColor: FAILURE_COLOR,
        "&:hover": {
            backgroundColor: FAILURE_COLOR
        }
    },
    buttonFailureMain: {
        backgroundColor: FAILURE_COLOR,
        "&:hover": {
            backgroundColor: FAILURE_COLOR
        },
        height: "120px"
    },
    buttonSuccessMain: {
        backgroundColor: SUCCESS_COLOR,
        "&:hover": {
            backgroundColor: SUCCESS_COLOR
        },
        height: "120px"
    },
    fabProgress: {
        color: LOADING_COLOR,
        position: "absolute",
        top: -6,
        left: -6,
        zIndex: 1
    },
    buttonProgress: {
        color: LOADING_COLOR,
        position: "absolute",
        top: "50%",
        left: "50%",
        marginTop: -12,
        marginLeft: -12
    }
});

export { styles, theme };
