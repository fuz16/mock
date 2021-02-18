package mock

import "sync"

var (
	locker  = sync.Mutex{}
	patches = make(map[uintptr]*patch)
)

func applyWithLock(p *patch) {
	locker.Lock()
	defer locker.Unlock()
	if pp, exist := patches[p.src]; exist {
		pp.undo()
	}

	p.apply()
	patches[p.src] = p
}

func undoWithLock(ptr uintptr) {
	locker.Lock()
	defer locker.Unlock()
	if pp, exist := patches[ptr]; exist {
		pp.undo()
		return
	}
	delete(patches, ptr)
}

func undoAll() {
	locker.Lock()
	defer locker.Unlock()
	for _, p := range patches {
		p.undo()
		delete(patches, p.src)
	}
}
