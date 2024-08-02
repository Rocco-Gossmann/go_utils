package go_utils_test

import (
	"fmt"
	"testing"

	"github.com/rocco-gossmann/go_utils"
	"github.com/rocco-gossmann/go_utils/testdata"
	"github.com/stretchr/testify/assert"
)

func TestCopyWithProgress(t *testing.T) {

	t.Run("copy 4 bytes with progress", func(t *testing.T) {

		var testWriter = testdata.NewTestWriter(t)
		var testReader = testdata.NewTestReader(t)

		var output string

		var onProgress = func(progress int) {
			output = fmt.Sprintf("%s%d", output, progress)
		}

		var err error
		var copied int64 = 0

		copied, err = go_utils.CopyWithProgress(testReader, testWriter, onProgress)

		assert.Nil(t, err, "did not expect to see an error")
		assert.Equal(t, int64(4), copied, "expected 4 bytes to be copied")
		assert.Equal(t, "1234", output, "progress output should be '1234' but is not")

	})

}
