

Section
  - Tabs
  - Page
    - Main Content
    - Sidebar
  - Footer

SectionBase
- Init
  - initSection
    - Parse Configuration
    - Display Initial Message
  - CurrentPage.Init()
- Update
  - Cases
    - initialize
    - KeyMsgs
    - WindowResize
  - SyncProgramContext
  - Update Base
  - Update Tabs
  - Update Footer
  - CurrentPage.Update()
- View
  - Wrapper
    - Tabs
    - CurrentPage.View()
    - Footer

SectionImplementation
- Init
  - BaseModel.InitBase()
- Update
  - BaseModel.UpdateBase()
- View
  - BaseModel.ViewBase()

==========================

PageBase
- Init
  - initPage
    - Display Initial Message
  - CurrentPage.Init()
- Update
  - Cases
    - initialize
    - KeyMsgs
    - WindowResize
  - SyncProgramContext
  - Update Base
  - Update Main Content
  - Update Sidebar
  - CurrentPage.Update()
- View
  - Wrapper
    - Tabs
    - CurrentPage.View()
    - Footer

