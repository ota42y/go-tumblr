package tumblr

import (
  "testing"
)

func TestNew(t *testing.T) {
  actual := New()
  if actual == nil {
    t.Errorf("got %v\nwant %v", actual, nil)
  }
}
