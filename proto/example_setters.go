// Code generated by protoc-gen-go-setters. DO NOT EDIT.
// source: proto/example.proto
package protoc_gen_go

func (x *SearchRequest) SetQuery(in string) {
	x.Query = in
}

func (x *SearchRequest) SetPageNumber(in int32) {
	x.PageNumber = in
}

func (x *SearchRequest) SetResultsPerPage(in int32) {
	x.ResultsPerPage = in
}

func (x *SearchRequest) SetCorpus(in Corpus) {
	x.Corpus = in
}

func (x *SearchRequest) SetActive(in bool) {
	x.Active = in
}

func (x *SearchRequest) SetTags(in []string) {
	x.Tags = in
}

func (x *SearchRequest) AppendTags(in ...string) {
	x.Tags = append(x.Tags, in...)
}

func (x *SearchRequest) SetRegex(in *string) {
	x.Regex = in
}

func (x *HelloRequest) SetName(in string) {
	x.Name = in
}

func (x *HelloReply) SetMessage(in string) {
	x.Message = in
}
