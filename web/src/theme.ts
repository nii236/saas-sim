import { Theme, ThemeProvider } from "theme-ui"
import { SystemCssProperties } from "@styled-system/css"
import { flex } from "styled-system"
import { bulma } from "@theme-ui/presets"

interface AppTheme extends Theme {
	layout: { [key: string]: SystemCssProperties }
}
// export const theme = bulma

export const theme: AppTheme = {
	...bulma,
	sizes: ["0", ".25rem", ".5rem", "1rem", "2rem", "4rem", "8rem", "16rem"],
	space: ["0rem", "1rem", "2rem", "4rem", "8rem", "16rem", "32rem", "64rem"],
	// fontSizes: [12, 14, 16, 20, 24, 32],
	breakpoints: ["40em", "52em", "64em"],
}
