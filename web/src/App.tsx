/** @jsx jsx */
import { jsx, Layout, SxProps } from "theme-ui"
import { Styled } from "theme-ui"
import { Flex, Text, ITheme, Slider, SliderTrack, SliderFilledTrack, SliderThumb, Stack, Box } from "@chakra-ui/core"
import * as React from "react"
import "react-vis/dist/style.css"
import { XAxis, YAxis, FlexibleWidthXYPlot, LineSeries } from "react-vis"
import { Header, Main, Container, Footer } from "theme-ui"
interface SliderProps {
	value: number
}

import { TypeStyle, TypeScale, useTheme, HeadingStyle, FontFamily, ColorSwatch, ColorPalette, ThemeCard } from "@theme-ui/style-guide"
import { color, flex } from "styled-system"
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
		<Flex
			color="white"
			sx={{
				height: "100%",
				flexDirection: "column",
			}}>
			<Flex bg={"black"}>
				<Text
					as="h1"
					sx={{
						padding: 0,
						margin: 0,
					}}>
					Header
				</Text>
			</Flex>
			<Flex
				w={"100%"}
				sx={{
					padding: 0,
					flexGrow: 1,
					flexDirection: "row",
				}}>
				<Box w="7">
					<Flex flexDirection={"column"} width="100%" padding={2} bg={"black"} height="100%">
						<Text as="h2">Params</Text>
						<Stack spacing={"0px"}>
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
				<Box w="100%">
					<Flex
						sx={{
							padding: 0,
						}}
						flexDirection={"column"}>
						<h1>Results</h1>
						{/* <LineChart data={data} /> */}
						<Box>
							<StyleGuide sx={{ color: "black" }} />
						</Box>
					</Flex>
				</Box>
			</Flex>
			<Flex
				sx={{
					padding: 0,
					margin: [0],
				}}
				bg={"black"}>
				<Text>Footer</Text>
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
