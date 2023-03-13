package main

import (
	testmodel "groupie/testmodels"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var myApp = app.New()
var myWindow = myApp.NewWindow("Groupie Tracker")
var grid = container.New(layout.NewGridLayoutWithColumns(5))
var screen = container.NewHBox()
var leftVerticalBox = container.NewVBox()
var rightVerticalBox = container.NewVBox()
var rightColumnGroup = container.NewGridWithColumns(1)
var twoColumnsGroup = container.NewGridWithColumns(2)
var infoGroup = container.NewVBox()
var month string
var buttonStatus = 0
var buttonStatus2 = 0
var i = 1

func main() {
	pageMain()
	myWindow.CenterOnScreen()
	myApp.Settings().SetTheme(theme.DarkTheme())
	myWindow.Padded()
	myWindow.ShowAndRun()
}

func pageMain() {
	navbr := navbar()
	grid.Add(navbr)
	// pour executer une seule fois cette fonction
	if i == 1 {
		for id := 1; id < 53; id++ {
			grid.Add(card(id))
		}
		i = 2
	}
	screen := container.NewVScroll(grid)
	myWindow.Resize(fyne.NewSize(1450, 750))
	myWindow.SetContent(screen)
}

func card(id int) *fyne.Container {
	r, _ := fyne.LoadResourceFromURLString(testmodel.GetArtistsID(id).Image)
	img := canvas.NewImageFromResource(r)
	img.FillMode = canvas.ImageFillOriginal
	// img.SetMinSize(fyne.NewSize(230, 230))
	btn := widget.NewButton(testmodel.GetArtistsID(id).Name, func() {
		GetInfosBtn(id)
	})
	container1 := container.New(
		layout.NewVBoxLayout(),
		img,
		btn,
	)
	return container1
}
func navbar() *fyne.Container {
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Rechercher...")

	searchButton := widget.NewButton("Rechercher", func() {
		for i := 0; i < 54; i++ {
			if strings.EqualFold(searchEntry.Text, testmodel.GetArtistsID(i).Name) {
				GetInfosBtn(i)
			}
		}
	})
	lightThemeBtn := widget.NewButton("Light mode", func() {
		myApp.Settings().SetTheme(theme.LightTheme())
	})
	darkThemeBtn := widget.NewButton("Dark mode", func() {
		myApp.Settings().SetTheme(theme.DarkTheme())
	})
	container1 := container.New(
		layout.NewVBoxLayout(),
		searchEntry,
		searchButton,
		darkThemeBtn,
		lightThemeBtn,
	)
	return container1
}
func GetInfosBtn(id int) {
	rightColumnGroup.RemoveAll()
	buttonStatus = 0
	r, _ := fyne.LoadResourceFromURLString(testmodel.GetArtistsID(id).Image)
	img := canvas.NewImageFromResource(r)
	img.FillMode = canvas.ImageFillOriginal
	img.SetMinSize(fyne.NewSize(185, 185))
	contain := pageInfos(id)
	homeBtn := home()
	locationBtn := handleLocation(id)
	relationBtn := relationButton(id)
	leftVerticalBox = container.NewVBox(img, locationBtn, relationBtn, homeBtn)
	rightVerticalBox = container.NewVBox(contain, rightColumnGroup)
	screen = container.NewHBox(leftVerticalBox, rightVerticalBox)
	myWindow.SetContent(screen)
}

func pageInfos(id int) *fyne.Container {
	var listmember string
	var listmember2 string
	infoGroup = container.NewVBox()
	i := 0
	v := 0
	nameLabel := "Nom du groupe : " + testmodel.GetArtistsID(id).Name
	nameText := widget.NewLabel(nameLabel)
	nameText.Alignment = fyne.TextAlignLeading
	infoGroup.Add(nameText)
	members := testmodel.GetArtistsID(id).Members
	numberOfMembers := len(members)
	for _, member := range members {
		if i < (numberOfMembers / 2) {
			listmember = listmember + member + ", "
		} else if i == (numberOfMembers - 1) {
			listmember2 = listmember2 + member + ". "
			v = 1
		} else if i == (numberOfMembers - 2) {
			listmember2 = listmember2 + member + " et "
		} else {
			listmember2 = listmember2 + member + ", "
			v = 1
		}
		i = i + 1
	}
	memberLabel := "Membres du groupe : " + listmember
	memberText := widget.NewLabel(memberLabel)
	memberText.Alignment = fyne.TextAlignLeading
	infoGroup.Add(memberText)
	if v == 1 {
		stringmember2 := listmember2
		text5 := widget.NewLabel(stringmember2)
		text5.Alignment = fyne.TextAlignLeading
		infoGroup.Add(text5)
	}
	creationdateLabel := "Date de creation du groupe : " + strconv.Itoa(testmodel.GetArtistsID(id).CreationDate)
	creationdateText := widget.NewLabel(creationdateLabel)
	creationdateText.Alignment = fyne.TextAlignLeading
	infoGroup.Add(creationdateText)
	FirstAlbumLabel := "Date de sortie premier album : " + formatDate(testmodel.GetArtistsID(id).FirstAlbum)
	FirstAlbumText := widget.NewLabel(FirstAlbumLabel)
	FirstAlbumText.Alignment = fyne.TextAlignLeading
	infoGroup.Add(FirstAlbumText)
	return infoGroup
}

func home() *fyne.Container {
	retour := widget.NewButton("Retour à l'Accueil", func() {
		pageMain()
	})
	contain := container.NewVBox(retour)
	return contain
}

func handleLocation(id int) *fyne.Container {
	locationbtn := widget.NewButton("Locations / Dates", func() {
		// fmt.Print("LocationButton")
		getLocation(id)
	})
	contain := container.NewVBox(locationbtn)
	return contain
}

func relationButton(id int) *fyne.Container {
	relationbtn := widget.NewButton("Relation", func() {
		// fmt.Print("relationBtn")
		getRelations(id)
	})
	contain := container.NewVBox(relationbtn)
	return contain
}

func getRelations(id int) {
	rightColumnGroup.RemoveAll()
	buttonStatus = 0
	var locationAndDate string
	var title string
	var formattedDate string
	var formattedDates []string
	var container = container.NewGridWithColumns(1)
	title = "Concerts : ,"
	for location, dates := range testmodel.GetRelations(id).DatesLocations {
		location = cases.Title(language.Und).String(location)
		if len(dates) > 1 {
			for _, rightVerticalBox := range dates {
				datetemp := formatDate(rightVerticalBox)
				formattedDates = append(formattedDates, datetemp)
			}
			formattedDate = strings.Join(formattedDates, ", le ")
		} else {
			formattedDate = formatDate(strings.Join(dates, " "))
		}
		locationl := strings.Split(location, "_")
		locationli := strings.Join(locationl, " ")
		locationAndDate = "- " + cases.Title(language.Und).String(locationli) + " : le " + formattedDate + ","
		title = title + locationAndDate
	}
	var listOfDates []string = strings.Split(title, ",")

	reverse(listOfDates)
	buttonStatus = 0
	for _, val := range listOfDates {
		contain := widget.NewLabel(val)
		if buttonStatus == 0 {
			goBackBtn := widget.NewButton("Retour", func() {
				rightColumnGroup.RemoveAll()
				twoColumnsGroup.RemoveAll()
				buttonStatus = 0
				buttonStatus2 = 0
			})
			twoColumnsGroup.Add(goBackBtn)
			rightColumnGroup.Add(twoColumnsGroup)
			container.Add(contain)
			buttonStatus = 1
		} else if buttonStatus == 1 {
			rightColumnGroup.Remove(twoColumnsGroup)
			twoColumnsGroup.RemoveAll()
			buttonStatus = 0
		}

		// buttonStatus2 = 1
	}
	infoGroup.Add(container)
}

func getLocation(id int) {
	rightColumnGroup.RemoveAll()
	buttonStatus = 0
	if buttonStatus == 0 {
		var loca string
		var list string
		var container = container.NewGridWithColumns(1)
		fr := 0
		for _, place := range testmodel.GetLocations(id).Location {
			locationl := strings.Split(place, "_")
			locationli := strings.Join(locationl, " ")
			loca = cases.Title(language.Und).String(locationli)
			fr = fr + 1
			if fr == len(testmodel.GetLocations(id).Location) {
				list = list + loca
			} else {
				list = list + loca + ","
			}
		}
		var splittedLocation []string = strings.Split(list, ",")
		for _, locations := range splittedLocation {
			place := locations
			contain := widget.NewButton(place,
				func() {
					// fmt.Print(place)
					getDateInLocation(id, place)
				},
			)
			container.Add(contain)
		}
		rightColumnGroup.Add(container)
		buttonStatus = 1
	} else if buttonStatus == 1 {
		rightColumnGroup.RemoveAll()
		buttonStatus = 0
	}
}

func getDateInLocation(id int, locate string) {
	rightColumnGroup.RemoveAll()
	buttonStatus2 = 0
	if buttonStatus2 == 0 {
		rightColumnGroup.RemoveAll()
		text := "En concert à " + locate + " : "
		contain := widget.NewLabel(text)
		twoColumnsGroup.Add(contain)
		contain = widget.NewLabel(" ")
		twoColumnsGroup.Add(contain)
		varlocation := strings.Split(locate, " ")
		locationString := strings.Join(varlocation, "_")
		lowerCase := strings.ToLower(locationString)
		for location, dates := range testmodel.GetRelations(id).DatesLocations {
			if lowerCase == location {
				for _, rightVerticalBox := range dates {
					datetemp := formatDate(rightVerticalBox)
					datetemp = " -" + datetemp
					contain := widget.NewLabel(datetemp)
					contain.Alignment = fyne.TextAlignLeading
					contain.Alignment = fyne.TextAlign(fyne.TextWrapWord)
					twoColumnsGroup.Add(contain)
				}
			}
		}
		goBackBtn := widget.NewButton("Retour", func() {
			rightColumnGroup.RemoveAll()
			twoColumnsGroup.RemoveAll()
			buttonStatus2 = 0
			buttonStatus = 0
			getLocation(id)
		})
		twoColumnsGroup.Add(goBackBtn)
		rightColumnGroup.Add(twoColumnsGroup)
		buttonStatus2 = 1
	} else if buttonStatus2 == 1 {
		rightColumnGroup.Remove(twoColumnsGroup)
		twoColumnsGroup.RemoveAll()
		buttonStatus2 = 0
	}
}

func formatDate(date string) string {
	splittedDate := strings.Split(date, "-")
	switch splittedDate[1] {
	case "01":
		month = "janvier"
	case "02":
		month = "février"
	case "03":
		month = "mars"
	case "04":
		month = "avril"
	case "05":
		month = "mai"
	case "06":
		month = "juin"
	case "07":
		month = "juillet"
	case "08":
		month = "août"
	case "09":
		month = "septembre"
	case "10":
		month = "octobre"
	case "11":
		month = "novembre"
	case "12":
		month = "décembre"
	}
	splittedDate[1] = month
	joinedDate := strings.Join(splittedDate, " ")
	return joinedDate
}

func reverse(ss []string) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}
