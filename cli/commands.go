package cli

import "github.com/tucnak/climax"

var List = climax.Command{
	Name:  "list",
	Brief: "list all feature flags",
	Usage: `[-p=] "<prefix>" list all flags with a matching prefix`,
	Help:  `Lists all feature flags. Use -p to match flags by a prefix.`,

	Flags: []climax.Flag{
		{
			Name:     "prefix",
			Short:    "p",
			Usage:    `--prefix="<flag_name>"`,
			Help:     `List only flags with matching prefix.`,
			Variable: true,
		},
	},

	Examples: []climax.Example{
		{
			Usecase:     `-p "flag_"`,
			Description: `Matches 'flag_name'`,
		},
	},

	Handle: func(ctx climax.Context) int {
		return 0
	},
}

var Set = climax.Command{
	Name:  "set",
	Brief: "create or update a feature flag",
	Usage: `set -name flag_name -type [boolean|percentile] -value [0.0-1.0|true/false] -comment "flag description"`,
	Help:  `set creates or updates a feature flag.`,

	Flags: []climax.Flag{
		{
			Name:     "name",
			Short:    "n",
			Usage:    `--name="flag_name"`,
			Help:     `the name of the falg to set`,
			Variable: true,
		},
		{
			Name:     "type",
			Short:    "t",
			Usage:    `--type=[boolean|percentile]`,
			Help:     `the type of flag to set`,
			Variable: true,
		},
		{
			Name:     "value",
			Short:    "v",
			Usage:    `--value=0.0-1.0 or true|false`,
			Help:     `the value of the flag`,
			Variable: true,
		},
		{
			Name:     "comment",
			Short:    "c",
			Usage:    `--comment="flag description"`,
			Help:     `an optional comment or description`,
			Variable: true,
		},
	},

	Examples: []climax.Example{
		{
			Usecase:     `-n "flag_name" -t percentile -v 0.5 -c "the flag desc"`,
			Description: `sets a percentile flag to 50%`,
		},
		{
			Usecase:     `-n "flag_name" -t boolean -v false -c "the flag desc"`,
			Description: `sets a boolean flag to false`,
		},
	},

	Handle: func(ctx climax.Context) int {
		return 0
	},
}

var Delete = climax.Command{
	Name:  "delete",
	Brief: "delete a feature flag",
	Usage: `[-n=] "<name>" delete flag with matching name`,
	Help:  `Delete a feature flag matching --name`,

	Flags: []climax.Flag{
		{
			Name:     "name",
			Short:    "n",
			Usage:    `--name="<flag_name>"`,
			Help:     `Name of the flag to delete`,
			Variable: true,
		},
	},

	Examples: []climax.Example{
		{
			Usecase:     `-n "flag_name"`,
			Description: `Deletes 'flag_name'`,
		},
	},

	Handle: func(ctx climax.Context) int {
		return 0
	},
}

var Init = climax.Command{
	Name:  "init",
	Brief: "initialize the audit repo",
	Usage: ``,
	Help:  `Initiializes a new audit repo and pushes to the origin.`,

	Handle: func(ctx climax.Context) int {
		return 0
	},
}
