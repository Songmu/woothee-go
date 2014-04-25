/* This file is auto-generated! Any changes to this file will be lost! */
package woothee

import (
	"testing"
)

func Test_appliance(t *testing.T) {
	var result *Result
	var err error

	result, err = Parse(`Mozilla/5.0 (Nintendo 3DS; U; ; ja) Version/1.7455.JP`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Nintendo 3DS; U; ; ja) Version/1.7455.JP': %s`, err)
	} else {
		if result.Category != "appliance" {
			t.Errorf("Expected result.Category to be 'appliance', but got '%s'", result.Category)
		}
		if result.Os != "Nintendo 3DS" {
			t.Errorf("Expected result.Os to be 'Nintendo 3DS', but got '%s'", result.Os)
		}
		if result.Name != "Nintendo 3DS" {
			t.Errorf("Expected result.Name to be 'Nintendo 3DS', but got '%s'", result.Name)
		}
	}
	result, err = Parse(`Opera/9.50 (Nintendo DSi; Opera/507; U; ja)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Opera/9.50 (Nintendo DSi; Opera/507; U; ja)': %s`, err)
	} else {
		if result.Name != "Opera" {
			t.Errorf("Expected result.Name to be 'Opera', but got '%s'", result.Name)
		}
		if result.Version != "9.50" {
			t.Errorf("Expected result.Version to be '9.50', but got '%s'", result.Version)
		}
		if result.Os != "Nintendo DSi" {
			t.Errorf("Expected result.Os to be 'Nintendo DSi', but got '%s'", result.Os)
		}
		if result.Category != "appliance" {
			t.Errorf("Expected result.Category to be 'appliance', but got '%s'", result.Category)
		}
	}
	result, err = Parse(`Opera/9.30 (Nintendo Wii; U; ; 3642; ja)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Opera/9.30 (Nintendo Wii; U; ; 3642; ja)': %s`, err)
	} else {
		if result.Category != "appliance" {
			t.Errorf("Expected result.Category to be 'appliance', but got '%s'", result.Category)
		}
		if result.Os != "Nintendo Wii" {
			t.Errorf("Expected result.Os to be 'Nintendo Wii', but got '%s'", result.Os)
		}
		if result.Name != "Opera" {
			t.Errorf("Expected result.Name to be 'Opera', but got '%s'", result.Name)
		}
		if result.Version != "9.30" {
			t.Errorf("Expected result.Version to be '9.30', but got '%s'", result.Version)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Nintendo WiiU) AppleWebKit/534.52 (KHTML, like Gecko) NX/2.1.0.8.21 NintendoBrowser/1.0.0.7494.US`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Nintendo WiiU) AppleWebKit/534.52 (KHTML, like Gecko) NX/2.1.0.8.21 NintendoBrowser/1.0.0.7494.US': %s`, err)
	} else {
		if result.Category != "appliance" {
			t.Errorf("Expected result.Category to be 'appliance', but got '%s'", result.Category)
		}
		if result.Os != "Nintendo Wii U" {
			t.Errorf("Expected result.Os to be 'Nintendo Wii U', but got '%s'", result.Os)
		}
		if result.Name != "Nintendo Wii U" {
			t.Errorf("Expected result.Name to be 'Nintendo Wii U', but got '%s'", result.Name)
		}
	}
	result, err = Parse(`Mozilla/5.0 (PLAYSTATION 3; 1.00)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (PLAYSTATION 3; 1.00)': %s`, err)
	} else {
		if result.Name != "PlayStation 3" {
			t.Errorf("Expected result.Name to be 'PlayStation 3', but got '%s'", result.Name)
		}
		if result.Os != "PlayStation 3" {
			t.Errorf("Expected result.Os to be 'PlayStation 3', but got '%s'", result.Os)
		}
		if result.Category != "appliance" {
			t.Errorf("Expected result.Category to be 'appliance', but got '%s'", result.Category)
		}
	}
	result, err = Parse(`Mozilla/5.0 (PLAYSTATION 3 4.31) AppleWebKit/531.22.8 (KHTML, like Gecko)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (PLAYSTATION 3 4.31) AppleWebKit/531.22.8 (KHTML, like Gecko)': %s`, err)
	} else {
		if result.Category != "appliance" {
			t.Errorf("Expected result.Category to be 'appliance', but got '%s'", result.Category)
		}
		if result.Os != "PlayStation 3" {
			t.Errorf("Expected result.Os to be 'PlayStation 3', but got '%s'", result.Os)
		}
		if result.Name != "PlayStation 3" {
			t.Errorf("Expected result.Name to be 'PlayStation 3', but got '%s'", result.Name)
		}
	}
	result, err = Parse(`Mozilla/5.0 (PlayStation 4 1.000) AppleWebKit/536.26 (KHTML, like Gecko)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (PlayStation 4 1.000) AppleWebKit/536.26 (KHTML, like Gecko)': %s`, err)
	} else {
		if result.Category != "appliance" {
			t.Errorf("Expected result.Category to be 'appliance', but got '%s'", result.Category)
		}
		if result.Os != "PlayStation 4" {
			t.Errorf("Expected result.Os to be 'PlayStation 4', but got '%s'", result.Os)
		}
		if result.Name != "PlayStation 4" {
			t.Errorf("Expected result.Name to be 'PlayStation 4', but got '%s'", result.Name)
		}
	}
	result, err = Parse(`Mozilla/4.0 (PSP (PlayStation Portable); 2.00)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/4.0 (PSP (PlayStation Portable); 2.00)': %s`, err)
	} else {
		if result.Category != "appliance" {
			t.Errorf("Expected result.Category to be 'appliance', but got '%s'", result.Category)
		}
		if result.Os != "PlayStation Portable" {
			t.Errorf("Expected result.Os to be 'PlayStation Portable', but got '%s'", result.Os)
		}
		if result.Name != "PlayStation Portable" {
			t.Errorf("Expected result.Name to be 'PlayStation Portable', but got '%s'", result.Name)
		}
	}
	result, err = Parse(`Mozilla/5.0 (PlayStation Vita 1.51) AppleWebKit/531.22.8 (KHTML, like Gecko) Silk/3.2`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (PlayStation Vita 1.51) AppleWebKit/531.22.8 (KHTML, like Gecko) Silk/3.2': %s`, err)
	} else {
		if result.Os != "PlayStation Vita" {
			t.Errorf("Expected result.Os to be 'PlayStation Vita', but got '%s'", result.Os)
		}
		if result.Category != "appliance" {
			t.Errorf("Expected result.Category to be 'appliance', but got '%s'", result.Category)
		}
		if result.Name != "PlayStation Vita" {
			t.Errorf("Expected result.Name to be 'PlayStation Vita', but got '%s'", result.Name)
		}
	}
	result, err = Parse(`Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0; Xbox)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0; Xbox)': %s`, err)
	} else {
		if result.Name != "Xbox 360" {
			t.Errorf("Expected result.Name to be 'Xbox 360', but got '%s'", result.Name)
		}
		if result.Os != "Xbox 360" {
			t.Errorf("Expected result.Os to be 'Xbox 360', but got '%s'", result.Os)
		}
		if result.Category != "appliance" {
			t.Errorf("Expected result.Category to be 'appliance', but got '%s'", result.Category)
		}
	}
	result, err = Parse(`Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0; Xbox; Xbox One)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0; Xbox; Xbox One)': %s`, err)
	} else {
		if result.Name != "Xbox One" {
			t.Errorf("Expected result.Name to be 'Xbox One', but got '%s'", result.Name)
		}
		if result.Category != "appliance" {
			t.Errorf("Expected result.Category to be 'appliance', but got '%s'", result.Category)
		}
		if result.Os != "Xbox One" {
			t.Errorf("Expected result.Os to be 'Xbox One', but got '%s'", result.Os)
		}
	}
	result, err = Parse(`Mozilla/5.0 (DTV; TSBNetTV/TXXXXXXXXX.0203.ADD; like Gecko) NetFront/3.4 DTVNetBrowser/2.2 (000039;TXXXXXXXXX;0203;ADD) InettvBrowser/2.2 (000039;TXXXXXXXXX;0203;ADD) YahooDTV/1.1 (Video=0x03;Audio=0x01;Screen=0x01;Device=0x00;Remote=0x10)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (DTV; TSBNetTV/TXXXXXXXXX.0203.ADD; like Gecko) NetFront/3.4 DTVNetBrowser/2.2 (000039;TXXXXXXXXX;0203;ADD) InettvBrowser/2.2 (000039;TXXXXXXXXX;0203;ADD) YahooDTV/1.1 (Video=0x03;Audio=0x01;Screen=0x01;Device=0x00;Remote=0x10)': %s`, err)
	} else {
		if result.Name != "InternetTVBrowser" {
			t.Errorf("Expected result.Name to be 'InternetTVBrowser', but got '%s'", result.Name)
		}
		if result.Category != "appliance" {
			t.Errorf("Expected result.Category to be 'appliance', but got '%s'", result.Category)
		}
		if result.Os != "DigitalTV" {
			t.Errorf("Expected result.Os to be 'DigitalTV', but got '%s'", result.Os)
		}
	}
	result, err = Parse(`Mozilla/5.0 (DTV; TVwithVideoPlayer) NetFront/4.1 InettvBrowser/2.2 (08001F;DTV04VSFC3;0001;0001)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (DTV; TVwithVideoPlayer) NetFront/4.1 InettvBrowser/2.2 (08001F;DTV04VSFC3;0001;0001)': %s`, err)
	} else {
		if result.Name != "InternetTVBrowser" {
			t.Errorf("Expected result.Name to be 'InternetTVBrowser', but got '%s'", result.Name)
		}
		if result.Category != "appliance" {
			t.Errorf("Expected result.Category to be 'appliance', but got '%s'", result.Category)
		}
		if result.Os != "DigitalTV" {
			t.Errorf("Expected result.Os to be 'DigitalTV', but got '%s'", result.Os)
		}
	}
	result, err = Parse(`Mozilla/5.0 (DTV; TSBNetTV/T45000006.0203.CDD; like Gecko) NetFront/3.4 DTVNetBrowser/2.2 (000039;T45011C06;0203;CDD) InettvBrowser/2.2 (000039;T45011C06;0203;CDD)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (DTV; TSBNetTV/T45000006.0203.CDD; like Gecko) NetFront/3.4 DTVNetBrowser/2.2 (000039;T45011C06;0203;CDD) InettvBrowser/2.2 (000039;T45011C06;0203;CDD)': %s`, err)
	} else {
		if result.Name != "InternetTVBrowser" {
			t.Errorf("Expected result.Name to be 'InternetTVBrowser', but got '%s'", result.Name)
		}
		if result.Category != "appliance" {
			t.Errorf("Expected result.Category to be 'appliance', but got '%s'", result.Category)
		}
		if result.Os != "DigitalTV" {
			t.Errorf("Expected result.Os to be 'DigitalTV', but got '%s'", result.Os)
		}
	}
	result, err = Parse(`Mozilla/5.0 (Standard; NF34SW/1.1; like Gecko) NetFront/3.4 InettvBrowser/2.2C (000087;IP03-01;0100;0000)`)
	if err != nil {
		t.Errorf(`Failed to parse 'Mozilla/5.0 (Standard; NF34SW/1.1; like Gecko) NetFront/3.4 InettvBrowser/2.2C (000087;IP03-01;0100;0000)': %s`, err)
	} else {
		if result.Category != "appliance" {
			t.Errorf("Expected result.Category to be 'appliance', but got '%s'", result.Category)
		}
		if result.Os != "DigitalTV" {
			t.Errorf("Expected result.Os to be 'DigitalTV', but got '%s'", result.Os)
		}
		if result.Name != "InternetTVBrowser" {
			t.Errorf("Expected result.Name to be 'InternetTVBrowser', but got '%s'", result.Name)
		}
	}
}
