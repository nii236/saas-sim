/** @jsx jsx */
import { jsx, SxProps } from "theme-ui"
import { Flex, Text, Slider, SliderTrack, SliderFilledTrack, SliderThumb, Stack, Box } from "@chakra-ui/core"
import * as React from "react"
import "react-vis/dist/style.css"
import { XAxis, YAxis, FlexibleWidthXYPlot, LineSeries } from "react-vis"

import { TypeStyle, TypeScale, useTheme, HeadingStyle, FontFamily, ColorSwatch, ColorPalette, ThemeCard } from "@theme-ui/style-guide"
import { AppTheme } from "./theme"

export const StyleGuide = (props: { sx: SxProps["sx"] }) => {
	const theme: AppTheme = useTheme()

	const swatches = []
	for (const color in theme.colors) {
		console.log(color)
		swatches.push(
			<Box height={6} width={6} bg={color}>
				<Text textAlign="center">{color}</Text>
			</Box>,
		)
	}
	return (
		<div style={{ color: "black" }} sx={props.sx}>
			<h1>Style Guide</h1>
			<Flex wrap={"wrap"}>{swatches}</Flex>

			{/* <TypeScale /> */}
			{/* <TypeStyle fontFamily="heading" fontWeight="heading" lineHeight="heading" />
			<TypeStyle fontFamily="body" fontWeight="body" lineHeight="body" /> */}
			{/* <pre>{JSON.stringify(theme, null, 2)}</pre> */}
		</div>
	)
}

export const AppHeader = () => {
	return (
		<Box bg={"black"}>
			<Text
				as="h1"
				sx={{
					padding: 0,
					margin: 0,
				}}>
				Header
			</Text>
		</Box>
	)
}
export const AppSidebar = () => {
	return (
		<Box w={250}>
			<Flex direction={"column"} w="100%" p={2} bg={"black"} h="100%">
				<Text as="h2">Params</Text>
				<Stack spacing={4}>
					<Slider paddingY={[1]} defaultValue={30}>
						<SliderTrack />
						<SliderFilledTrack bg="primary" />
						<SliderThumb />
					</Slider>
					<Slider paddingY={[1]} defaultValue={30}>
						<SliderTrack />
						<SliderFilledTrack bg="primary" />
						<SliderThumb />
					</Slider>
					<Slider paddingY={[1]} defaultValue={30}>
						<SliderTrack />
						<SliderFilledTrack bg="primary" />
						<SliderThumb />
					</Slider>
				</Stack>

				<button
					sx={{
						padding: 0,
					}}>
					Simulate
				</button>
			</Flex>
		</Box>
	)
}
export const AppMainbar: React.FC = props => {
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
		<Box color="black">
			<h1>Results</h1>
			<LineChart data={data} />
			<StyleGuide sx={{ color: "black" }} />
			Lorem ipsum dolor sit amet consectetur adipisicing elit. Numquam necessitatibus deserunt aspernatur, pariatur illum beatae illo exercitationem aperiam
			quas? Dicta sequi sapiente illum tempore odio velit error dolore cumque rem?
		</Box>
	)
}

export const AppFooter = () => {
	return (
		<Box
			sx={{
				padding: 0,
				margin: [0],
			}}
			bg={"black"}>
			<Text>Footer</Text>
		</Box>
	)
}
export const App = () => {
	return (
		<Flex direction={"column"}>
			<Flex direction={"row"}>
				<AppHeader />
			</Flex>
			<Flex flexGrow={1} direction={"row"}>
				<AppSidebar />
				<AppMainbar />
			</Flex>
			<Flex direction={"row"}>
				<AppFooter />
			</Flex>
		</Flex>
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
