package main

type WorldConfig struct {
	CreateParam    ParamCreateWorld `json:"create_param"`
	RunParam       ParamRunWorld    `json:"run_param"`
	Bodies         []*Body          `json:"bodies"`
	StaticBodyIds  []string         `json:"static_body_ids"`
	DynamicBodyIds []string         `json:"dynamic_body_ids"`
}
