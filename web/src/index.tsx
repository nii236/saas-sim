import * as React from "react"
import * as ReactDOM from "react-dom"
import { App } from "./App"
import "./styles.scss"
import "tachyons"
import { ThemeProvider } from "theme-ui"
import { theme } from "./theme"

var mountNode = document.getElementById("app")
ReactDOM.render(
	<ThemeProvider theme={theme}>
		<App />
	</ThemeProvider>,
	mountNode,
)
