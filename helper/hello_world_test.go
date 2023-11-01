package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Rahmat",
			request: "Rahmat",
		},
		{
			name:    "Budi",
			request: "Budi",
		},
		{
			name:    "Joko",
			request: "Joko",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Rahmat", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Rahmat")
		}
	})

	b.Run("Budi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Budi")
		}
	})
}

func BenchmarkHelloWorldRahmat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Rahmat")
	}
}

func BenchmarkHelloWorldBudi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Budi")
	}
}

func TestTableHelloWorld(t *testing.T) {
	// tests data
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Rahmat",
			request:  "Rahmat",
			expected: "Hello Rahmat",
		},
		{
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},
		{
			name:     "Joko",
			request:  "Joko",
			expected: "Hello Joko",
		},
		{
			name:     "Dodi",
			request:  "Dodi",
			expected: "Hello Dodi",
		},
	}

	// iteration test
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}

}

func TestSubTest(t *testing.T) {
	// sub test 1
	t.Run("Rahmat", func(t *testing.T) {
		var result string = HelloWorld("Rahmat")
		require.Equal(t, "Hello Rahmat", result, "Result must be 'Hello Rahmat'")
	})

	// sub test 2
	t.Run("Budi", func(t *testing.T) {
		var result string = HelloWorld("Budi")
		require.Equal(t, "Hello Budi", result, "Result must be 'Hello Budi'")
	})
}

func TestMain(m *testing.M) {
	// before test run
	fmt.Println("BEFORE TEST RUN")

	// test run
	m.Run()

	// after test run
	fmt.Println("AFTER TEST RUN")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Cannot run on Windows OS")
	}
	var result string = HelloWorld("Rahmat")
	require.Equal(t, "Hello Rahmat", result, "Result must be 'Hello Rahmat'")
}

func TestHelloWorldRequire(t *testing.T) {
	var result string = HelloWorld("Rahmat")
	require.Equal(t, "Hello Rahmat", result, "Result must be 'Hello Rahmat'")
	fmt.Println("TestHelloWorld with Require Done") // this line just for check is execution line after or not
}

func TestHelloWorldAssert(t *testing.T) {
	var result string = HelloWorld("Rahmat")
	assert.Equal(t, "Hello Rahmat", result, "Result must be 'Hello Rahmat'")
	fmt.Println("TestHelloWorld with Assert Done") // this line just for check is execution line after or not
}

func TestHelloWorldRahmat(t *testing.T) {
	var result string = HelloWorld("Rahmat")
	if result != "Hello Rahmat" {
		// panic("Result is not 'Hello Rahmat'") // dont use panic to return fail
		// t.Fail() // use Fail not show what's the error?
		t.Error("Result is not 'Hello Rahmat'")
	}
	fmt.Println("TestHelloWorldRahmat Done") // this line just for check is execution line after or not
}

func TestHelloWorldBudi(t *testing.T) {
	var result string = HelloWorld("Budi")
	if result != "Hello Budi" {
		// panic("Result is not 'Hello Budi'") // dont use panic to return fail
		// t.FailNow() // same like Fail but not execution new line program and not show the error
		t.Fatal("Result is not 'Hello Budi'")
	}
	fmt.Println("TestHelloWorldBudi Done") // this line just for check is execution line after or not
}
