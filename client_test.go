package quickwit

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	ctx := context.Background()

	// Setup Quickwit testcontainer
	container, err := SetupQuickwitContainer(ctx)
	require.NoError(t, err)

	t.Cleanup(func() {
		err := container.Terminate(ctx)
		if err != nil {
			t.Logf("Failed to terminate container: %v", err)
		}
	})

	// Create client with testcontainer endpoint
	client := New(WithEndpoint(container.Endpoint))

	t.Run("Get Cluster", func(t *testing.T) {
		cluster, err := client.GetCluster(ctx)
		require.NoError(t, err)
		assert.NotNil(t, cluster)
		assert.NotEmpty(t, cluster.ClusterID)

		t.Logf("Cluster: %+v", cluster)
	})

	t.Run("Create Index", func(t *testing.T) {
		indexConfig := IndexConfig{
			Version: "0.9",
			ID:      "test-index",
			DocMapping: DocMapping{
				Mode: "dynamic",
				FieldMappings: []FieldMapping{
					{
						Name:    "timestamp",
						Type:    "datetime",
						Indexed: true,
						Stored:  true,
					},
					{
						Name:    "message",
						Type:    "text",
						Indexed: true,
						Stored:  true,
					},
				},
			},
			IndexingSettings: Settings{
				Resources: Resources{
					HeapSize: "100MB",
				},
			},
			SearchSettings: SearchSettings{
				DefaultSearchFields: []string{"message"},
			},
		}

		idx, err := client.CreateIndex(ctx, indexConfig)
		require.NoError(t, err)
		assert.NotNil(t, idx)
		assert.Equal(t, "test-index", idx.Config.ID)
	})

	t.Run("List Indexes", func(t *testing.T) {
		// List indexes and verify our index is present
		indexes, err := client.ListIndexes(ctx)
		require.NoError(t, err)
		assert.NotNil(t, indexes)
		assert.NotEmpty(t, indexes)

		// Verify our created index is in the list
		found := false
		for _, idx := range indexes {
			if idx.Config.ID == "test-index" {
				found = true
				break
			}
		}
		assert.True(t, found, "Expected to find 'list-test-index' in the list of indexes")
	})

	t.Run("Get Index", func(t *testing.T) {
		// Get the index
		idx, err := client.GetIndex(ctx, "test-index")
		require.NoError(t, err)
		assert.NotNil(t, idx)
		assert.Equal(t, "test-index", idx.Config.ID)
	})

	t.Run("Delete Index", func(t *testing.T) {
		// Create an index first
		indexConfig := IndexConfig{
			Version: "0.9",
			ID:      "delete-test-index",
			DocMapping: DocMapping{
				Mode: "dynamic",
				FieldMappings: []FieldMapping{
					{
						Name:    "body",
						Type:    "text",
						Indexed: true,
						Stored:  true,
					},
				},
			},
			IndexingSettings: Settings{
				Resources: Resources{
					HeapSize: "100MB",
				},
			},
			SearchSettings: SearchSettings{
				DefaultSearchFields: []string{"body"},
			},
		}

		_, err := client.CreateIndex(ctx, indexConfig)
		require.NoError(t, err)

		// Delete the index
		err = client.DeleteIndex(ctx, "delete-test-index")
		require.NoError(t, err)

		// Verify it's deleted
		_, err = client.GetIndex(ctx, "delete-test-index")
		assert.Error(t, err)
	})

	t.Run("Describe Index", func(t *testing.T) {
		// Describe the index
		desc, err := client.DescribeIndex(ctx, "test-index")
		require.NoError(t, err)
		assert.NotNil(t, desc)
		assert.Equal(t, "test-index", desc.IndexID)
	})

	t.Run("Clear Index", func(t *testing.T) {
		// Clear the index
		err = client.ClearIndex(ctx, "test-index")
		require.NoError(t, err)
	})

	t.Run("List Splits", func(t *testing.T) {
		// List splits
		splits, err := client.ListSplits(ctx, "test-index")
		require.NoError(t, err)
		assert.NotNil(t, splits)
		assert.IsType(t, &SplitsRes{}, splits)
	})

	t.Run("Create and Delete Source", func(t *testing.T) {
		// Skip this test as file sources require local CLI usage
		// Quickwit doesn't allow file sources via API for security reasons
		t.Skip("File sources are not allowed via API, use CLI command 'quickwit tool local-ingest'")
	})

	t.Run("Search", func(t *testing.T) {
		// Search on the test-index (empty index, should return no results but no error)
		result, err := client.Search(ctx, "test-index", "*")
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 0, result.NumHits)
	})

	t.Run("Get Elastic Endpoint", func(t *testing.T) {
		cluster, err := client.GetElastic(ctx)
		require.NoError(t, err)
		assert.NotNil(t, cluster)
		// The Elastic endpoint may return a different structure
		// Just verify we get a valid response
	})
}
