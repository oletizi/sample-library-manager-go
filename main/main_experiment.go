package main

/*
func mainExperiment() {

	ds := samplelib.NewFilesystemDataSource("test/data/library/multi-level")

	node, err := ds.RootNode()
	if err != nil {
		log.Fatal(err)
	}
	// Create the list to display the node contents
	nodeView := tview.NewList()
	nodeView.SetBorder(true).SetTitle(" Node View ")
	nodeView.ShowSecondaryText(false)
	nodeView.SetHighlightFullLine(true)

	// Create the info view
	infoView := tview.NewTextView()
	infoView.SetBorder(true).SetTitle(" Info View ")
	infoView.Clear()

	// create the log view
	logView := tview.NewTextView()
	logView.SetBorder(true).SetTitle(" Log ")
	logView.Clear()
	logger := &logger{view: logView}

	flex := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(nodeView, 0, 1, true).
		AddItem(infoView, 0, 1, true).
		AddItem(logView, 0, 1, true)

	// Set up the model (or is this the controller? I guess it's sort of both)
	model := &appModel{logger: logger, ds: ds, node: node, nodeView: nodeView, infoView: infoView}
	nodeView.SetChangedFunc(model.OnNodeViewChange)
	nodeView.SetSelectedFunc(model.OnNodeViewSelected)
	model.UpdateNodeView()

	app := tview.NewApplication()
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Name() == "Ctrl+L" {
			logView.Clear()
		}
		logger.Println(fmt.Sprintf("Event name: %s, key: %d, rune: %d", event.Name(), event.Key(), event.Rune()))
		return event
	})
	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}

type appModel struct {
	logger   Logger
	ds       samplelib.DataSource
	node     samplelib.Node
	nodes    []samplelib.Node
	samples  []samplelib.Sample
	nodeView *tview.List
	infoView *tview.TextView
	updating bool
}

// UpdateNodeView populates the node view with the contents of the current node
func (m *appModel) UpdateNodeView() *appModel {
	m.logger.Println("Turn on updating flag")
	// *all* changes, including structural changes, get sent to OnNodeViewChange
	// this flag tells OnNodeViewChange to ignore changes while it's being updated.
	m.updating = true
	m.nodes = nil
	m.samples = nil
	m.nodeView.Clear()
	m.nodeView.SetTitle(" " + m.node.Name() + " ")
	if !m.node.Parent().Null() {
		m.logger.Println("Adding an item for the parent")
		parent := m.node.Parent()
		// Add item has weird behavior, so using InsertItem instead
		m.nodes = append(m.nodes, parent)
		m.nodeView.AddItem("..", "", 0, nil)
	}
	ds := m.ds

	m.logger.Println(fmt.Sprintf("Getting children of node: %s", m.node.Name()))
	children, err := ds.ChildrenOf(m.node)
	if err != nil {
		m.logger.Println(err)
	}
	m.logger.Println(fmt.Sprintf("Child count: %d", len(children)))
	for _, child := range children {
		m.logger.Println("Adding list item for child node: " + child.Path())
		m.nodes = append(m.nodes, child)
		m.nodeView.AddItem(child.Name()+"/", "", 0, nil)
	}

	samples, err := ds.SamplesOf(m.node)
	if err != nil {
		return nil
	}
	for _, sample := range samples {
		m.samples = append(m.samples, sample)
		m.nodeView.AddItem(sample.Name(), "", 0, nil)
	}
	m.logger.Println("Turn off updating flag")
	m.updating = false
	return m
}

// OnNodeViewChange is the callback function for when anything in the node view (tview list) changes.
// It figures out what to put in the info panel.
func (m *appModel) OnNodeViewChange(index int, _ string, _ string, _ rune) {
	m.logger.Println("list changed!")
	if m.updating {
		m.logger.Println("The list is updating, so we'll ignore this change.")
		return
	}
	// Figure out what to put in the info panel
	info := ""
	nodes, err := m.ds.ChildrenOf(m.node)
	if err != nil {
		m.logger.Println(err)
	}
	if !m.node.Parent().Null() {
		nodes = append(nodes, m.node.Parent())
	}
	if index < len(nodes) {
		// we're on a node, so show the node info in the Info View
		node := nodes[index]
		info += fmt.Sprintf("Path: %s\nName: %s", node.Path(), node.Name())
	} else {
		// we're on a sample, so show the sample info in the Info View
		samples, err := m.ds.SamplesOf(m.node)
		if err != nil {
			m.logger.Println(err)
		}
		sampleIndex := index - len(nodes)
		m.logger.Println(fmt.Sprintf("Sample index: %d; sample len: %d", sampleIndex, len(samples)))
		//sample := samples[sampleIndex]
		//info += fmt.Sprintf("Path: %s\nName: %s", sample.Path, sample.Name)
	}

	// update the info view
	m.logger.Println(info)
	m.infoView.Clear()
	_, err = m.infoView.Write([]byte(info))
	if err != nil {
		m.logger.Println(err)
	}
}

func (m *appModel) OnNodeViewSelected(i int, _ string, _ string, _ rune) {
	if i < len(m.nodes) {
		// this is a node
		m.logger.Println("We selected a node: " + m.nodes[i].Name())
		m.node = m.nodes[i]
		m.UpdateNodeView()
	} else if i < len(m.nodes)+len(m.samples) {
		// this is a sample
		m.logger.Println("We selected a sample: " + m.samples[i-len(m.nodes)].Name())
	} else {
		m.logger.Println(fmt.Sprintf("YIKES! We selected something out of bounds; i: %d, nodes: %d, samples: %d",
			i, len(m.nodes), len(m.samples)))
	}
}

type Logger interface {
	Println(msg any) Logger
}
type logger struct {
	view *tview.TextView
}

func (l *logger) Println(msg any) Logger {
	w := l.view.BatchWriter()
	defer func(w tview.TextViewWriter) {
		err := w.Close()
		if err != nil {
			log.Print(err)
		}
	}(w)
	_, err := fmt.Fprintln(w, msg)
	if err != nil {
		log.Print(err)
	}
	return l
}
*/
