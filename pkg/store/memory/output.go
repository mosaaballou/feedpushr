package store

import (
	"github.com/ncarlier/feedpushr/pkg/common"
	"github.com/ncarlier/feedpushr/pkg/model"
)

func (store *InMemoryStore) nextOutputSequence() int {
	max := 0
	for _, output := range store.outputs {
		if output.ID > max {
			max = output.ID
		}
	}
	return max + 1
}

// ClearOutputs clear all outputs
func (store *InMemoryStore) ClearOutputs() error {
	store.outputsLock.RLock()
	defer store.outputsLock.RUnlock()
	store.outputs = make(map[int]model.OutputDef)
	return nil
}

// GetOutput returns a stored Output.
func (store *InMemoryStore) GetOutput(ID int) (*model.OutputDef, error) {
	output, exists := store.outputs[ID]
	if !exists {
		return nil, common.ErrOutputNotFound
	}
	return &output, nil
}

// DeleteOutput removes a output.
func (store *InMemoryStore) DeleteOutput(ID int) (*model.OutputDef, error) {
	store.outputsLock.RLock()
	defer store.outputsLock.RUnlock()
	output, err := store.GetOutput(ID)
	if err != nil {
		return nil, err
	}
	delete(store.outputs, ID)
	return output, nil
}

// SaveOutput stores a output.
func (store *InMemoryStore) SaveOutput(output model.OutputDef) (*model.OutputDef, error) {
	store.outputsLock.RLock()
	defer store.outputsLock.RUnlock()
	if output.ID == 0 {
		output.ID = store.nextOutputSequence()
	}
	store.outputs[output.ID] = output
	return &output, nil
}

// ListOutputs returns a paginated list of outputs.
func (store *InMemoryStore) ListOutputs(page, limit int) (*model.OutputDefCollection, error) {
	outputs := model.OutputDefCollection{}
	startOffset := (page - 1) * limit
	offset := 0
	for _, output := range store.outputs {
		switch {
		case offset < startOffset:
			// Skip entries before the start offset
			offset++
			continue
		case offset >= startOffset+limit:
			// End of the window
			break
		default:
			// Add value to entries
			offset++
			outputs = append(outputs, &output)
		}
	}
	return &outputs, nil
}

// ForEachOutput iterates over all outputs
func (store *InMemoryStore) ForEachOutput(cb func(*model.OutputDef) error) error {
	store.outputsLock.RLock()
	defer store.outputsLock.RUnlock()
	for _, output := range store.outputs {
		if err := cb(&output); err != nil {
			return err
		}
	}
	return nil
}
