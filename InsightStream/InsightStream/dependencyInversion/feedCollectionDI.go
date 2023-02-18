package dependencyInversion

import "insightstream/ent"

type (
	FeedCollection interface {
		FetchAll(cl *ent.Client) error
		FetchSingle() error
	}

	FeedCollectionImpl struct {
	}
)

func NewFeedCollection() FeedCollection {
	return &FeedCollectionImpl{}
}

func (f *FeedCollectionImpl) FetchAll(cl *ent.Client) error {

	return nil

}

func (f *FeedCollectionImpl) FetchSingle() error {
	return nil
}
