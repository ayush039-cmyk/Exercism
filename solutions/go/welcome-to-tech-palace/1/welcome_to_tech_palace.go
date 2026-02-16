package techpalace
import "strings"
import "fmt"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    l := strings.ToUpper(customer)
	x:= fmt.Sprintf("Welcome to the Tech Palace, %s",l)
    return x
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := strings.Repeat("*", numStarsPerLine)
    return border + "\n" + welcomeMsg + "\n" + border
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    noStars := strings.ReplaceAll(oldMsg, "*", "")
    cleaned := strings.TrimSpace(noStars)
    
    return cleaned
}
