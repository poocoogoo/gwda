package gwda

import (
	"fmt"
	"testing"
)

func TestElement_Click(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_ = s.AppLaunch(bundleId)
	Debug = true
	element, err := s.FindElement(WDALocator{LinkText: NewWDAElementAttribute().SetValue("通知")})
	if err != nil {
		// staleElementReferenceErrorWithMessage
		t.Fatal(err)
	}
	t.Log(element)

	err = element.Click()
	if err != nil {
		t.Fatal(err)
	}
}

func TestElement_SendKeys(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_ = s.AppLaunch(bundleId)
	Debug = true
	element, err := s.FindElement(WDALocator{ClassName: WDAElementType{SearchField: true}})
	if err != nil {
		t.Fatal(err)
	}

	err = element.SendKeys(bundleId)
	if err != nil {
		t.Fatal(err)
	}
}

func TestElement_Clear(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	defer s.DeleteSession()
	_ = s.AppLaunch(bundleId)
	Debug = true
	element, err := s.FindElement(WDALocator{ClassName: WDAElementType{SearchField: true}})
	if err != nil {
		t.Fatal(err)
	}

	err = element.SendKeys(bundleId)
	if err != nil {
		t.Fatal(err)
	}

	err = element.Clear()
	if err != nil {
		t.Fatal(err)
	}
}

func TestElement_Rect(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	element, err := s.FindElement(WDALocator{LinkText: NewWDAElementAttribute().SetValue("通知")})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(element)

	rect, err := element.Rect()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rect)
}

func TestElement_IsEnabled(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	bundleId := "com.apple.Preferences"
	_ = bundleId
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	element, err := s.FindElement(WDALocator{LinkText: NewWDAElementAttribute().SetValue("通知")})
	if err != nil {
		t.Fatal(err)
	}

	isEnabled, err := element.IsEnabled()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(isEnabled)
}

func TestElement_IsDisplayed(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_ = s.AppLaunch(bundleId)
	Debug = true
	element, err := s.FindElement(WDALocator{LinkText: NewWDAElementAttribute().SetValue("通知")})
	if err != nil {
		t.Fatal(err)
	}

	displayed, err := element.IsDisplayed()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(displayed)
}

func TestElement_IsSelected(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_ = s.AppLaunch(bundleId)
	Debug = true
	element, err := s.FindElement(WDALocator{LinkText: NewWDAElementAttribute().SetValue("通知")})
	if err != nil {
		t.Fatal(err)
	}

	isSelected, err := element.IsSelected()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(isSelected)

	if isSelected {
		return
	}

	// iPad 左右分栏
	element, err = s.FindElement(WDALocator{Predicate: "selected == true AND label == '通用'"})
	if err != nil {
		t.Fatal(err)
	}
	isSelected, err = element.IsSelected()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(isSelected)
}

func TestElement_GetAttribute(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	element, err := s.FindElement(WDALocator{LinkText: NewWDAElementAttribute().SetLabel("通用")})
	if err != nil {
		t.Fatal(err)
	}

	// attrName := "type"
	attr := NewWDAElementAttribute().SetUID("")
	// attr = NewWDAElementAttribute().SetAccessibilityContainer(false)
	value, err := element.GetAttribute(attr)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(attr.getAttributeName(), "=", value)
	t.Log("element.UID", "=", element.UID)
}

func TestElement_Text(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_ = s.AppLaunch(bundleId)
	Debug = true
	element, err := s.FindElement(WDALocator{LinkText: NewWDAElementAttribute().SetLabel("通用")})
	if err != nil {
		t.Fatal(err)
	}

	// attrName := "type"
	// attrName = NewWDAElementAttribute().SetUID("").GetAttributeName()
	// attrName = NewWDAElementAttribute().SetAccessibilityContainer(false).GetAttributeName()
	text, err := element.Text()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(text)
}

func TestElement_Type(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_ = s.AppLaunch(bundleId)
	Debug = true
	element, err := s.FindElement(WDALocator{LinkText: NewWDAElementAttribute().SetLabel("通用")})
	if err != nil {
		t.Fatal(err)
	}

	elemType, err := element.Type()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(elemType)
	t.Log(elemType == WDAElementType{Cell: true}.String())
	t.Log(elemType == fmt.Sprintf("%s", WDAElementType{StaticText: true}))
}

func TestElement_Screenshot(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_ = s.AppLaunch(bundleId)
	Debug = true
	element, err := s.FindElement(WDALocator{Predicate: "type == 'XCUIElementTypeCell' AND name == '通知'"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(element)

	_, err = element.Screenshot()
	if err != nil {
		t.Fatal(err)
	}
}

func TestElement_ScreenshotToJpeg(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_ = s.AppLaunch(bundleId)
	Debug = true
	// element, err := s.FindElement(WDALocator{LinkText: NewWDAElementAttribute().SetValue("通知")})
	element, err := s.FindElement(WDALocator{Predicate: "type == 'XCUIElementTypeCell' AND name == '通知'"})
	if err != nil {
		t.Fatal(err)
	}

	toPng, err := element.ScreenshotToJpeg()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("元素图片的大小", toPng.Bounds().Size())
	t.Log(element.Rect())
}

func TestElement_ScreenshotToDiskAsJpeg(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_ = s.AppLaunch(bundleId)
	Debug = true
	// element, err := s.FindElement(WDALocator{LinkText: NewWDAElementAttribute().SetValue("通知")})
	element, err := s.FindElement(WDALocator{Predicate: "type == 'XCUIElementTypeCell' AND name == '通知'"})
	if err != nil {
		t.Fatal(err)
	}

	err = element.ScreenshotToDiskAsJpeg("/Users/hero/Desktop/e1.png")
	if err != nil {
		t.Fatal(err)
	}
}

func TestElement_Tmp(t *testing.T) {

	// parse, err2 := url.Parse("http://localhost:8100/session/1A742166-05E8-467D-8A6C-3C775FD9AEAD")
	// if err2 != nil {
	// 	t.Fatal(err2)
	// }
	// elem := newElement(parse, "21000000-0000-0000-960D-000000000000")
	// // [FBRoute GET:@"/wda/element/:uuid/accessible"]
	//
	// fmt.Println(strings.Repeat("-", 100))
	//
	// fmt.Println(elem._withFormat("/accessible"))
	// fmt.Println(urlJoin(elem.endpoint, elem._withFormat("/accessible")))
	//
	// fmt.Println(strings.Repeat("*", 50))
	//
	// tmp, _ := url.Parse(elem.endpoint.String())
	// tmp.Path = path.Join(elem.endpoint.Path, "wda", elem._withFormat("/accessible"))
	// fmt.Println(tmp.String())
	//
	// fmt.Println(strings.Repeat("*", 50))
	//
	// fmt.Println(urlJoin(elem.endpoint, elem._withFormat("/accessible"), true))
	// fmt.Println(urlJoin(elem.endpoint, elem._withFormat("/accessible"), false))
	//
	// fmt.Println(strings.Repeat("-", 100))
	//
	// return

	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	element, err := s.FindElement(WDALocator{LinkText: NewWDAElementAttribute().SetLabel("通用")})
	// element, err := s.FindElement(WDALocator{Predicate: "selected == true AND label == '通用'"})
	// element, err := s.FindElements(WDALocator{LinkText: NewWDAElementAttribute().SetLabel("通用")})
	if err != nil {
		t.Fatal(err)
	}

	// addToRootWda(element.elementURL)
	// return

	element.tttTmp()
	// t.Log(element.GetAttribute(NewWDAElementAttribute().SetUID("")))

	// for _, elem := range element {
	// 	elem.tttTmp()
	// 	t.Log(elem.GetAttribute(NewWDAElementAttribute().SetUID("")))
	// }

}
