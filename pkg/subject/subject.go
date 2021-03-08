package subject

// Subject created by an admin.
type Subject struct {
	ID        uint      `json:"id,omitempty"`
	Name 	  string	`json:"name,omitempty"`
	Year  	  int	    `json:"year,omitempty"`
}