package dependencyInversion

type (
	FeedCollection interface {
		FetchAll() error
		FetchSingle() error
	}

	FeedCollectionImpl struct {
	}
)

func NewFeedCollection() FeedCollection {
	return &FeedCollectionImpl{}
}

func (f *FeedCollectionImpl) FetchAll() error {

	return nil

}

func (f *FeedCollectionImpl) FetchSingle() error {
	return nil
}
