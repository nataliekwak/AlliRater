package newsfeed

import "testing"

//use package testing which allows for automated testing of go packages

//use error methods to indicate that the result was not what we intended it to be
func TestAdd(t *testing.T) {
	feed := New()
	//creates a new feed "object"?
	feed.Add(Item{})
	//adds empty slice of items

	//the length of the items slice should be 1, but if it is not then the item has not been added
	if len(feed.Items) != 1 {
		t.Errorf("Item was not added")
	}
}

//determines if the get all function works
//similarly we add an item to our feed

func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Item{})
	results := feed.GetAll()
	if len(results) != 1 {
		t.Errorf("Item was not added")
	}
}
