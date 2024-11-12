#!/bin/bash

set -e
set -u
set -o pipefail

read -p "Enter the Leetcode problem number: " problem_number
read -p "Enter the Leetcode problem slug: " problem_slug

directory_name="solutions/${problem_number}_${problem_slug}"

if [ ! -d "solutions" ]; then
    echo "The solutions directory does not exist."
    exit 1
fi

if [ -d "$directory_name" ]; then
    echo "The directory '$directory_name' already exists."
    exit 1
fi

mkdir -p "$directory_name"
solution_file="${directory_name}/${problem_slug}.go"
solution_test_file="${directory_name}/${problem_slug}_test.go"

echo "package ${problem_slug}" > $solution_file

cat <<EOL > $solution_test_file
package ${problem_slug}

import "testing"

func TestMethodName(t *testing.T) {
    cases := []struct {
        input int
        want int
    }{
        // Add test cases here
        {0, 0},
        {0, 0},
    }

    for _, c := range cases {
        got := methodName(c.input)
        if got != c.want {
            t.Errorf("methodName(%v) = %d; want %d", c.input, got, c.want)
        }
    }
}

func BenchmarkMethodName(b *testing.B) {
    var result int
    for i := 0; i < b.N; i++ {
        result = methodName(0)
    }
    _ = result
}
EOL

gofmt -w $solution_file
gofmt -w $solution_test_file

echo "Directory '$directory_name' created with necessary Go files."