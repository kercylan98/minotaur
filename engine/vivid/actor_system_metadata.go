package vivid

import (
	"github.com/kercylan98/minotaur/toolkit/collection"
	"strings"
)

const (
	actorSystemMetadataDelimiter                           = ":-:"
	actorSystemClusterName                                 = "ClusterName"
	actorSystemMetadataKeySupervisionStrategyProviderTable = "SupervisionStrategyProviderTable"
	actorSystemMetadataKeyActorProviderTable               = "ActorProviderTable"
	actorSystemMetadataKeyMailboxProviderTable             = "MailboxProviderTable"
	actorSystemMetadataKeyDispatcherProviderTable          = "DispatcherProviderTable"
	actorSystemMetadataKeyPersistenceStorageProviderTable  = "PersistenceStorageProviderTable"
)

type metadata map[string]string

func packActorSystemMetadata(system *ActorSystem) metadata {
	md := make(metadata)

	md[actorSystemClusterName] = system.Name()
	md[actorSystemMetadataKeySupervisionStrategyProviderTable] = strings.Join(collection.ConvertMapKeysToSlice(system.config.supervisionStrategyProviderTable), actorSystemMetadataDelimiter)
	md[actorSystemMetadataKeyActorProviderTable] = strings.Join(collection.ConvertMapKeysToSlice(system.config.actorProviderTable), actorSystemMetadataDelimiter)
	md[actorSystemMetadataKeyMailboxProviderTable] = strings.Join(collection.ConvertMapKeysToSlice(system.config.mailboxProviderTable), actorSystemMetadataDelimiter)
	md[actorSystemMetadataKeyDispatcherProviderTable] = strings.Join(collection.ConvertMapKeysToSlice(system.config.dispatcherProviderTable), actorSystemMetadataDelimiter)
	md[actorSystemMetadataKeyPersistenceStorageProviderTable] = strings.Join(collection.ConvertMapKeysToSlice(system.config.persistenceStorageProviderTable), actorSystemMetadataDelimiter)

	return md
}
