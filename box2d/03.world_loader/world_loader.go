package main

import "encoding/json"

type WorldConfig struct {
	CreateParam   ParamCreateWorld   `json:"create_param"`
	RunParam      ParamRunWorld      `json:"run_param"`
	StaticBodies  []*ParamCreateBody `json:"-"`
	DynamicBodies []*ParamCreateBody `json:"-"`
}

func (c *WorldConfig) UnmarshalJSON(data []byte) error {
	type Alias WorldConfig
	aux := &struct {
		Bodies         []*ParamCreateBody `json:"bodies"`
		StaticBodyIds  []string           `json:"static_body_ids"`
		DynamicBodyIds []string           `json:"dynamic_body_ids"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	// get static bodies
	c.StaticBodies = []*ParamCreateBody{}
	for _, id := range aux.StaticBodyIds {
		for _, item := range aux.Bodies {
			if item.ID == id {
				c.StaticBodies = append(c.StaticBodies, item)
			}
		}
	}

	// get dynamic bodies
	c.DynamicBodies = []*ParamCreateBody{}
	for _, id := range aux.DynamicBodyIds {
		for _, item := range aux.Bodies {
			if item.ID == id {
				c.DynamicBodies = append(c.DynamicBodies, item)
			}
		}
	}
	return nil
}

func (c *WorldConfig) MarshalJSON() ([]byte, error) {
	type Alias WorldConfig
	staticIds := []string{}
	dynamicIds := []string{}
	bodies := []*ParamCreateBody{}

	for _, item := range c.StaticBodies {
		staticIds = append(staticIds, item.ID)
		bodies = append(bodies, item)
	}

	for _, item := range c.DynamicBodies {
		dynamicIds = append(dynamicIds, item.ID)
		bodies = append(bodies, item)
	}

	return json.Marshal(&struct {
		Bodies         []*ParamCreateBody `json:"bodies"`
		StaticBodyIds  []string           `json:"static_body_ids"`
		DynamicBodyIds []string           `json:"dynamic_body_ids"`
		*Alias
	}{
		StaticBodyIds:  staticIds,
		DynamicBodyIds: dynamicIds,
		Bodies:         bodies,
		Alias:          (*Alias)(c),
	})
}

func (c *WorldConfig) BuildWorld(scene *Scene) *World {
	// physics
	world := CreateWorld(c.CreateParam)

	// add static bodies
	for _, bodyParam := range c.StaticBodies {
		body := CreateBody(*bodyParam)
		world.AddBody(body)
	}

	// add dynamic bodies
	for _, bodyParam := range c.DynamicBodies {
		body := CreateBody(*bodyParam)
		world.AddBody(body)
	}

	return world
}

func (c *WorldConfig) RunWorld(sceneId string) *Scene {
	scene := CreateScene(sceneId)
	world := c.BuildWorld(scene)

	// run world
	world.Run(c.RunParam, scene)
	return scene
}
