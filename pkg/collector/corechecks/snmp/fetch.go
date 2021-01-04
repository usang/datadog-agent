package snmp

import (
	"fmt"
	"github.com/DataDog/datadog-agent/pkg/util/log"
	"sort"
)

func fetchValues(session sessionAPI, config snmpConfig) (*snmpValues, error) {
	scalarResults, err := fetchScalarOidsByBatch(session, config.OidConfig.scalarOids, config.OidBatchSize)
	if err != nil {
		return &snmpValues{}, fmt.Errorf("SNMPGET error: %v", err)
	}

	oids := make(map[string]string)
	for _, value := range config.OidConfig.columnOids {
		oids[value] = value
	}
	columnResults, err := fetchColumnOids(session, oids, config.OidBatchSize)
	if err != nil {
		return &snmpValues{}, fmt.Errorf("SNMPBULK error: %v", err)
	}

	return &snmpValues{scalarResults, columnResults}, nil
}

func fetchScalarOidsOne(session sessionAPI, oids []string) (map[string]snmpValue, error) {
	// Get results
	log.Debugf("fetchScalarOidsByBatch() oids: %v", oids)
	results, err := session.Get(oids)
	log.Debugf("fetchColumnOids() results: %v", results)
	if err != nil {
		return nil, err
	}
	return resultToScalarValues(results), nil
}

func fetchScalarOidsByBatch(session sessionAPI, oids []string, oidBatchSize int) (map[string]snmpValue, error) {
	// Get results
	// TODO: Improve batching algorithm and make it more readable
	retValues := make(map[string]snmpValue)

	if oidBatchSize == 0 {
		// TODO: test me
		return nil, fmt.Errorf("oidBatchSize cannot be 0")
	}

	// Make matches
	for i := 0; i < len(oids); i += oidBatchSize {
		j := i + oidBatchSize
		if j > len(oids) {
			j = len(oids)
		}
		fetchOids := oids[i:j]

		results, err := fetchScalarOidsOne(session, fetchOids)
		if err != nil {
			return nil, err
		}
		for k, v := range results {
			retValues[k] = v
		}
	}
	return retValues, nil
}

func fetchColumnOids(session sessionAPI, oids map[string]string, oidBatchSize int) (map[string]map[string]snmpValue, error) {
	// Get results
	// TODO: Improve batching algorithm and make it more readable

	if oidBatchSize == 0 {
		// TODO: test me
		return nil, fmt.Errorf("oidBatchSize cannot be 0")
	}

	retValues := make(map[string]map[string]snmpValue)
	columnOids := make([]string, 0, len(oids))

	for k := range oids {
		columnOids = append(columnOids, k)
	}

	// Make matches
	for i := 0; i < len(oids); i += oidBatchSize {
		j := i + oidBatchSize
		if j > len(oids) {
			j = len(oids)
		}
		batchColumnOids := columnOids[i:j]
		fetchOids := make(map[string]string)
		for _, oid := range batchColumnOids {
			fetchOids[oid] = oids[oid]
		}

		results, err := fetchColumnOidsOneBatch(session, fetchOids)
		if err != nil {
			return nil, err
		}
		for columnOid, instanceOids := range results {
			if _, ok := retValues[columnOid]; ok {
				for oid, value := range instanceOids {
					retValues[columnOid][oid] = value
				}
			} else {
				retValues[columnOid] = instanceOids
			}
		}
	}
	return retValues, nil
}

func fetchColumnOidsOneBatch(session sessionAPI, oids map[string]string) (map[string]map[string]snmpValue, error) {
	// Returns map[columnOID]map[index]interface(float64 or string)
	// GetBulk results
	// TODO:
	//   - make batches
	//   - GetBulk loop to get all rows
	returnValues := make(map[string]map[string]snmpValue)
	if len(oids) == 0 {
		return returnValues, nil
	}
	curOids := oids
	for {
		log.Debugf("fetchColumnOids() curOids  : %v", curOids)
		var columnOids, bulkOids []string
		for k, v := range curOids {
			columnOids = append(columnOids, k)
			bulkOids = append(bulkOids, v)
		}
		// sorting columnOids and bulkOids to make them deterministic for testing purpose
		sort.Strings(columnOids)
		sort.Strings(bulkOids)
		results, err := session.GetBulk(bulkOids)
		log.Debugf("fetchColumnOids() results: %v", results)
		if err != nil {
			return nil, err
		}
		values, nextOids := resultToColumnValues(columnOids, results)
		for columnOid, columnValues := range values {
			for oid, value := range columnValues {
				if _, ok := returnValues[columnOid]; !ok {
					returnValues[columnOid] = make(map[string]snmpValue)
				}
				returnValues[columnOid][oid] = value
			}
		}
		if len(nextOids) == 0 {
			break
		}
		curOids = nextOids
	}
	return returnValues, nil
}
