package vivid

import "github.com/kercylan98/minotaur/engine/vivid/behavior"

type ActorBehavior = behavior.Behavior[ActorContext]

type ActorPerformance = behavior.Performance[ActorContext]

type ActorFunctionalPerformance = behavior.FunctionalPerformance[ActorContext]
