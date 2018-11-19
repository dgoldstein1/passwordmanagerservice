// CardMediaWithIcon.js

import Button from "@material-ui/core/Button";
import CheckIcon from "@material-ui/icons/Check";
import WarningIcon from "@material-ui/icons/Warning";
import CardMedia from "@material-ui/core/CardMedia";
import loadingGif from "../images/loading.gif";
import React from "react";

/**
 * card media which displays specific image depending on success value
 * @param {json} material-ui classes
 * @param {string} 'loading', 'success' or 'failure'
 **/

let CardMediaWithIcon = (classes, status) => {
    return (
        <CardMedia className={classes.cardMedia} src={loadingGif}>
            {status === "loading" && (
                <div style={{ display: "flex", justifyContent: "center" }}>
                    <img
                        alt={status}
                        height="100px"
                        width="100px"
                        style={{ objectFit: "contain" }}
                        src={loadingGif}
                    />
                </div>
            )}
            {status === "success" && (
                <div style={{ display: "flex", justifyContent: "center" }}>
                    <Button
                        className={classes.buttonSuccessMain}
                        size={"large"}
                        fullWidth={true}
                    >
                        <CheckIcon />
                    </Button>
                </div>
            )}
            {status === "failure" && (
                <div style={{ display: "flex", justifyContent: "center" }}>
                    <Button
                        className={classes.buttonFailureMain}
                        size={"large"}
                        fullWidth={true}
                        disableRipple={true}
                    >
                        <WarningIcon />
                    </Button>
                </div>
            )}
        </CardMedia>
    );
};

export default CardMediaWithIcon;
