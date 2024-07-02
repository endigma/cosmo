package integration

import (
	"github.com/stretchr/testify/require"
	"github.com/wundergraph/cosmo/router-tests/testenv"
	"testing"
)

func TestFeatureFlags(t *testing.T) {

	t.Parallel()

	t.Run("Base feature graph schema is served when no feature flag is enabled", func(t *testing.T) {

		t.Parallel()

		testenv.Run(t, &testenv.Config{}, func(t *testing.T, xEnv *testenv.Environment) {
			res := xEnv.MakeGraphQLRequestOK(testenv.GraphQLRequest{
				Query: `{ employees { id productCount } }`,
			})
			require.Empty(t, res.Response.Header.Get("X-Feature-Flag"))
			require.JSONEq(t, `{"errors":[{"message":"field: productCount not defined on type: Employee","path":["query","employees","productCount"]}],"data":null}`, res.Body)
		})
	})

	t.Run("Base feature graph schema is served when feature flag does not exist / header", func(t *testing.T) {

		t.Parallel()

		testenv.Run(t, &testenv.Config{}, func(t *testing.T, xEnv *testenv.Environment) {
			res := xEnv.MakeGraphQLRequestOK(testenv.GraphQLRequest{
				Query: `{ employees { id productCount } }`,
				Header: map[string][]string{
					"X-Feature-Flag": {"nonexistent"},
				},
			})
			require.Empty(t, res.Response.Header.Get("X-Feature-Flag"))
			require.JSONEq(t, `{"errors":[{"message":"field: productCount not defined on type: Employee","path":["query","employees","productCount"]}],"data":null}`, res.Body)
		})
	})

	t.Run("Base feature graph schema is served when feature flag does not exist / cookie", func(t *testing.T) {

		t.Parallel()

		testenv.Run(t, &testenv.Config{}, func(t *testing.T, xEnv *testenv.Environment) {
			res := xEnv.MakeGraphQLRequestOK(testenv.GraphQLRequest{
				Query: `{ employees { id productCount } }`,
				Header: map[string][]string{
					"Cookie": {"feature_flag=nonexistent"},
				},
			})
			require.Empty(t, res.Response.Header.Get("X-Feature-Flag"))
			require.JSONEq(t, `{"errors":[{"message":"field: productCount not defined on type: Employee","path":["query","employees","productCount"]}],"data":null}`, res.Body)
		})
	})

	t.Run("Should replace product feature graph when feature flag is sent over header", func(t *testing.T) {

		t.Parallel()

		testenv.Run(t, &testenv.Config{}, func(t *testing.T, xEnv *testenv.Environment) {
			res := xEnv.MakeGraphQLRequestOK(testenv.GraphQLRequest{
				Query: `{ employees { id productCount } }`,
				Header: map[string][]string{
					"X-Feature-Flag": {"myff"},
				},
			})
			require.Equal(t, res.Response.Header.Get("X-Feature-Flag"), "myff")
			require.JSONEq(t, `{"data":{"employees":[{"id":1,"productCount":5},{"id":2,"productCount":2},{"id":3,"productCount":2},{"id":4,"productCount":3},{"id":5,"productCount":2},{"id":7,"productCount":0},{"id":8,"productCount":2},{"id":10,"productCount":3},{"id":11,"productCount":1},{"id":12,"productCount":4}]}}`, res.Body)
		})
	})

	t.Run("Should replace product feature graph when feature flag is sent over cookie", func(t *testing.T) {

		t.Parallel()

		testenv.Run(t, &testenv.Config{}, func(t *testing.T, xEnv *testenv.Environment) {
			res := xEnv.MakeGraphQLRequestOK(testenv.GraphQLRequest{
				Query: `{ employees { id productCount } }`,
				Header: map[string][]string{
					"Cookie": {"feature_flag=myff"},
				},
			})
			require.Equal(t, res.Response.Header.Get("X-Feature-Flag"), "myff")
			require.JSONEq(t, `{"data":{"employees":[{"id":1,"productCount":5},{"id":2,"productCount":2},{"id":3,"productCount":2},{"id":4,"productCount":3},{"id":5,"productCount":2},{"id":7,"productCount":0},{"id":8,"productCount":2},{"id":10,"productCount":3},{"id":11,"productCount":1},{"id":12,"productCount":4}]}}`, res.Body)
		})
	})

}
