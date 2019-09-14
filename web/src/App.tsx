/** @jsx jsx */
import { jsx, Layout } from "theme-ui"
import { Styled } from "theme-ui"

import * as React from "react"
import "react-vis/dist/style.css"
import { XAxis, YAxis, FlexibleWidthXYPlot, LineSeries } from "react-vis"
import { Header, Main, Container, Footer } from "theme-ui"
interface SliderProps {
	value: number
}

import { TypeStyle, TypeScale, useTheme, HeadingStyle, FontFamily, ColorSwatch, ColorPalette, ThemeCard } from "@theme-ui/style-guide"
const Slider: React.FC<SliderProps & React.InputHTMLAttributes<HTMLInputElement>> = props => {
	return (
		<input min="1" max="100" value={props.value} type="range">
			{props.children}
		</input>
	)
}
export const StyleGuide = () => {
	const theme = useTheme()
	console.log(theme)
	return (
		<React.Fragment>
			<h1>Style Guide</h1>
			{/* <ColorPalette sx={{ padding: 0, margin: 0 }} /> */}
			<TypeScale />
			{/* <TypeStyle fontFamily="heading" fontWeight="heading" lineHeight="heading" />
			<TypeStyle fontFamily="body" fontWeight="body" lineHeight="body" /> */}
			<div>
				<pre>{JSON.stringify(theme, null, 2)}</pre>
			</div>
		</React.Fragment>
	)
}
export const App = () => {
	const data = [
		{ x: 0, y: 8 },
		{ x: 1, y: 5 },
		{ x: 2, y: 4 },
		{ x: 3, y: 9 },
		{ x: 4, y: 1 },
		{ x: 5, y: 7 },
		{ x: 6, y: 6 },
		{ x: 7, y: 3 },
		{ x: 8, y: 2 },
		{ x: 9, y: 0 },
	]
	return (
		<Layout
			sx={{
				height: "100%",
				display: "flex",
				flexDirection: "column",
			}}>
			<Header>
				<h1
					sx={{
						width: "100%",
						padding: 2,
						margin: 0,
						color: "white",
						bg: "dark",
						display: "flex",
					}}>
					Header
				</h1>
			</Header>
			<Main
				sx={{
					padding: 0,
					flexGrow: 1,
					display: "flex",
					flexDirection: "row",
				}}>
				<Container
					sx={{
						flexDirection: "column",
						margin: [0],
						padding: 1,
						color: "white",
						bg: "dark",
						display: ["none", "flex"],
						width: [0, 7],
						variant: "layout.sidebar",
					}}>
					<h2>Params</h2>
					<Slider />
					<Slider />
					<div sx={{ flexGrow: 1 }}>
						<Slider />
					</div>

					<button
						sx={{
							padding: [0, 1],
							alignSelf: "center",
							backgroundColor: "primary",
							borderWidth: 1,
							borderColor: "white",
						}}>
						Regenerate
					</button>
				</Container>
				<Container
					sx={{
						padding: [2],
						bg: "white",
						flexGrow: 1,
						variant: "layout.mainbar",
					}}>
					{/* <h1>Results</h1> */}
					<LineChart data={data} />
					{/* <StyleGuide /> */}
				</Container>
			</Main>
			<Footer
				sx={{
					display: "flex",
					padding: [1],
					margin: [0],
					bg: "dark",
					color: "white",
				}}>
				Footer
			</Footer>
		</Layout>
	)
}

interface LineChartProps {
	data: { x: number; y: number }[]
}
const LineChart = (props: LineChartProps) => {
	return (
		<FlexibleWidthXYPlot height={300}>
			<LineSeries data={props.data} />
			<XAxis />
			<YAxis />
		</FlexibleWidthXYPlot>
	)
}
