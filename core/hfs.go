package main

const staticFileApi = "https://model.cortexlabs.ai/"

func hfsAccessUrl(ih string) string {
	return staticFileApi + ih + ".tar"
}
