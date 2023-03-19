package fetching

// CelebrationResponse is the response of the CelebrationQuery.
type CelebrationResponse struct {
	id     string
	date   string
	name   string
	status string
	cType  string
}

// NewCelebrationResponse creates a new CelebrationResponse.
func NewCelebrationResponse(id, date, name, status, t string) CelebrationResponse {
	return CelebrationResponse{id: id, date: date, name: name, status: status, cType: t}
}

// ID returns the id of the celebration.
func (r CelebrationResponse) ID() string {
	return r.id
}

// Date returns the date of the celebration.
func (r CelebrationResponse) Date() string {
	return r.date
}

// Name returns the name of the celebration.
func (r CelebrationResponse) Name() string {
	return r.name
}

// Status returns the status of the celebration.
func (r CelebrationResponse) Status() string {
	return r.status
}

// Type returns the type of the celebration.
func (r CelebrationResponse) Type() string {
	return r.cType
}
