#import <Appkit/Appkit.h>
#import "tab_view_delegate.h"
#import "_cgo_export.h"

@implementation NSTabViewDelegateAdaptor


- (void)tabViewDidChangeNumberOfTabViewItems:(NSTabView*)tabView {
    tabViewDelegate_TabViewDidChangeNumberOfTabViewItems([self goID], tabView);
}

- (BOOL)tabView:(NSTabView*)tabView shouldSelectTabViewItem:(NSTabViewItem*)tabViewItem {
    bool result_ = tabViewDelegate_TabView_ShouldSelectTabViewItem([self goID], tabView, tabViewItem);
    return result_;
}

- (void)tabView:(NSTabView*)tabView willSelectTabViewItem:(NSTabViewItem*)tabViewItem {
    tabViewDelegate_TabView_WillSelectTabViewItem([self goID], tabView, tabViewItem);
}

- (void)tabView:(NSTabView*)tabView didSelectTabViewItem:(NSTabViewItem*)tabViewItem {
    tabViewDelegate_TabView_DidSelectTabViewItem([self goID], tabView, tabViewItem);
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return TabViewDelegate_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteTabViewDelegate([self goID]);
	[super dealloc];
}
@end

void* WrapTabViewDelegate(uintptr_t goID) {
    NSTabViewDelegateAdaptor* adaptor = [[NSTabViewDelegateAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
