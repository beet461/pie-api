package main

type Test struct {
	TestRespond string
}

type SignIn struct {
	Type     string
	Email    string
	Password string
	Id       string
}

type Colors struct {
	EmailMenu              string
	EmailBackground        string
	EmailMenuBoxes         string
	BottomNavBarBack       string
	BottomNavBarSelected   string
	BottomNavBarUnselected string
}

type Font struct {
	HeaderSize            int
	HeaderFont            string
	GeneralTextSize       int
	GeneralFont           string
	EmailPreviewTitleSize int
	EmailPreviewTitleFont string
}

type Customise struct {
	Account string
	Colors  Colors
	Font    Font
}

func newFont() Font {
	return Font{20, "assets/fonts/Typo_Round_Light_Demo.otf", 15, "", 18, "assets/fonts/VictorMono-Light.ttf"}
}

func newColor() Colors {
	return Colors{"0xFF607D8B", "0xFF292929", "0xFF7094AC", "0xFF90A4AE", "0xFF636363", "0xFF0053DB"}
}
