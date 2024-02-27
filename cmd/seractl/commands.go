package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/seracoin/serad/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.SeradMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.SeradMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.SeradMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.SeradMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.SeradMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.SeradMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.SeradMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.SeradMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.SeradMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.SeradMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.SeradMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.SeradMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.SeradMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.SeradMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.SeradMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.SeradMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.SeradMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.SeradMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.SeradMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.SeradMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.SeradMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.SeradMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.SeradMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.SeradMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.SeradMessage_BanRequest{}),
	reflect.TypeOf(protowire.SeradMessage_UnbanRequest{}),
}

type commandDescription struct {
	name       string
	parameters []*parameterDescription
	typeof     reflect.Type
}

type parameterDescription struct {
	name   string
	typeof reflect.Type
}

func commandDescriptions() []*commandDescription {
	commandDescriptions := make([]*commandDescription, len(commandTypes))

	for i, commandTypeWrapped := range commandTypes {
		commandType := unwrapCommandType(commandTypeWrapped)

		name := strings.TrimSuffix(commandType.Name(), "RequestMessage")
		numFields := commandType.NumField()

		var parameters []*parameterDescription
		for i := 0; i < numFields; i++ {
			field := commandType.Field(i)

			if !isFieldExported(field) {
				continue
			}

			parameters = append(parameters, &parameterDescription{
				name:   field.Name,
				typeof: field.Type,
			})
		}
		commandDescriptions[i] = &commandDescription{
			name:       name,
			parameters: parameters,
			typeof:     commandTypeWrapped,
		}
	}

	return commandDescriptions
}

func (cd *commandDescription) help() string {
	sb := &strings.Builder{}
	sb.WriteString(cd.name)
	for _, parameter := range cd.parameters {
		_, _ = fmt.Fprintf(sb, " [%s]", parameter.name)
	}
	return sb.String()
}
