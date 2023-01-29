package main

import "fmt"

type Project struct {
	Name  string
	Stars int
}

func main() {
	projects := []Project{
		{
			Name:  "https://github.com/docker/compose",
			Stars: 27600,
		},
		{
			Name:  "https://github.com/docker/awesome-compose",
			Stars: 20924,
		},
		{
			Name:  "https://github.com/docker/kitematic",
			Stars: 12268,
		},
		{
			Name:  "https://github.com/docker/labs",
			Stars: 11097,
		},
		{
			Name:  "https://github.com/docker/docker-bench-security",
			Stars: 8173,
		},
		{
			Name:  "https://github.com/docker/dockercraft",
			Stars: 6814,
		},
		{
			Name:  "https://github.com/docker/machine",
			Stars: 6553,
		},
		{
			Name:  "https://github.com/docker/docker-py",
			Stars: 6080,
		},
		{
			Name:  "https://github.com/docker/cli",
			Stars: 5571,
		},
		{
			Name:  "https://github.com/docker/docs",
			Stars: 3734,
		},
		{
			Name:  "https://github.com/docker/build-push-action",
			Stars: 3025,
		},
		{
			Name:  "https://github.com/docker/libchan",
			Stars: 2459,
		},
		{
			Name:  "https://github.com/docker/buildx",
			Stars: 2374,
		},
	}
	reps := make(map[string]int)
	for i := range projects {
		reps[projects[i].Name] = projects[i].Stars
	}
	for k, v := range reps {
		fmt.Println(k, v)
	}
}
