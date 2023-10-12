package seed

import (
	"time"

	"github.com/pfumarola/ld-01/database/models"
	"gorm.io/gorm"
)

func seedAuthors(db *gorm.DB) {
	db.Create(authors)
}

var authors = []models.Author{
	{
		FirstName:   "George",
		LastName:    "Orwell",
		DateOfBirth: time.Date(1903, time.June, 25, 0, 0, 0, 0, time.UTC),
		Nationality: "British",
		Biography:   "George Orwell, whose real name was Eric Arthur Blair, was a British writer known for his dystopian novels, including '1984' and 'Animal Farm.' He was born on June 25, 1903, in Motihari, India, and later moved to England. Orwell's works continue to be influential and are often studied for their critique of totalitarianism and social injustice.",
	},
	{
		FirstName:   "Harper",
		LastName:    "Lee",
		DateOfBirth: time.Date(1926, time.April, 28, 0, 0, 0, 0, time.UTC),
		Nationality: "American",
		Biography:   "Harper Lee was an American novelist best known for her Pulitzer Prize-winning novel, 'To Kill a Mockingbird,' published in 1960. She was born on April 28, 1926, in Monroeville, Alabama. Lee's novel is considered a classic of modern American literature and addresses issues of racism and injustice in the American South.",
	},
	{
		FirstName:   "F. Scott",
		LastName:    "Fitzgerald",
		DateOfBirth: time.Date(1896, time.September, 24, 0, 0, 0, 0, time.UTC),
		Nationality: "American",
		Biography:   "F. Scott Fitzgerald was an American novelist and short story writer known for his novel 'The Great Gatsby,' published in 1925. He was born on September 24, 1896, in St. Paul, Minnesota. Fitzgerald's works often explore themes of wealth, decadence, and the American Dream during the Roaring Twenties.",
	},
	{
		FirstName:   "J.D.",
		LastName:    "Salinger",
		DateOfBirth: time.Date(1919, time.January, 1, 0, 0, 0, 0, time.UTC),
		Nationality: "American",
		Biography:   "J.D. Salinger was an American author best known for his novel 'The Catcher in the Rye,' published in 1951. He was born on January 1, 1919, in New York City. Salinger's novel is a classic of American literature and explores the struggles of its protagonist, Holden Caulfield, in a post-war society.",
	},
	{
		FirstName:   "Oscar",
		LastName:    "Wilde",
		DateOfBirth: time.Date(1854, time.October, 16, 0, 0, 0, 0, time.UTC),
		Nationality: "Irish",
		Biography:   "Oscar Wilde was an Irish playwright, poet, and author known for his wit and flamboyant style. He was born on October 16, 1854, in Dublin, Ireland. Wilde's works include 'The Picture of Dorian Gray' and 'The Importance of Being Earnest.' His writing is celebrated for its humor and social commentary.",
	},
	{
		FirstName:   "Bram",
		LastName:    "Stoker",
		DateOfBirth: time.Date(1847, time.November, 8, 0, 0, 0, 0, time.UTC),
		Nationality: "Irish",
		Biography:   "Bram Stoker was an Irish author best known for his Gothic horror novel 'Dracula,' published in 1897. He was born on November 8, 1847, in Clontarf, Dublin, Ireland. Stoker's novel has become a classic of the vampire genre and has inspired numerous adaptations and interpretations.",
	},
	{
		FirstName:   "Aldous",
		LastName:    "Huxley",
		DateOfBirth: time.Date(1894, time.July, 26, 0, 0, 0, 0, time.UTC),
		Nationality: "British",
		Biography:   "Aldous Huxley was an English writer and philosopher known for his novel 'Brave New World,' published in 1932. He was born on July 26, 1894, in Godalming, Surrey, England. Huxley's work often explores the impact of technology and societal control on human civilization.",
	},
	{
		FirstName:   "Ray",
		LastName:    "Bradbury",
		DateOfBirth: time.Date(1920, time.August, 22, 0, 0, 0, 0, time.UTC),
		Nationality: "American",
		Biography:   "Ray Bradbury was an American author known for his science fiction and fantasy works. He was born on August 22, 1920, in Waukegan, Illinois. Bradbury's notable works include 'Fahrenheit 451' and 'The Martian Chronicles,' which have left a lasting impact on the genre.",
	},
}
