import React from "react";
import { Wizard, WizardStep } from "@clutch-sh/wizard";
const WelcomeStep = () => (React.createElement(WizardStep, { isLoading: false, error: undefined }, "Hello World!"));
const HelloWorld = ({ heading }) => {
    const dataLayout = {};
    return (React.createElement(Wizard, { dataLayout: dataLayout, heading: heading },
        React.createElement(WelcomeStep, { name: "Welcome" })));
};
export default HelloWorld;
