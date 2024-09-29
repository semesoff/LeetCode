package main

import "container/list"

type Node struct {
    freq int
    keys map[string]struct{}
}

type AllOne struct {
    keyFreq  map[string]int
    freqList *list.List
    freqMap  map[int]*list.Element
}

func Constructor() AllOne {
    return AllOne{
        keyFreq:  make(map[string]int),
        freqList: list.New(),
        freqMap:  make(map[int]*list.Element),
    }
}

func (this *AllOne) Inc(key string) {
    freq := this.keyFreq[key]
    this.keyFreq[key]++

    // Remove key from current frequency node
    if freq > 0 {
        this.removeKeyFromNode(key, freq)
    }

    // Add key to new frequency node
    this.addKeyToNode(key, freq+1)
}

func (this *AllOne) Dec(key string) {
    freq := this.keyFreq[key]
    if freq == 0 {
        return
    }

    // Remove key from current frequency node
    this.removeKeyFromNode(key, freq)

    if freq == 1 {
        // Remove the key from keyFreq if its count becomes 0
        delete(this.keyFreq, key)
    } else {
        // Decrease frequency
        this.keyFreq[key]--
        this.addKeyToNode(key, freq-1)
    }
}

func (this *AllOne) GetMaxKey() string {
    if this.freqList.Len() == 0 {
        return ""
    }
    maxNode := this.freqList.Back().Value.(*Node)
    for key := range maxNode.keys {
        return key
    }
    return ""
}

func (this *AllOne) GetMinKey() string {
    if this.freqList.Len() == 0 {
        return ""
    }
    minNode := this.freqList.Front().Value.(*Node)
    for key := range minNode.keys {
        return key
    }
    return ""
}

// Helper function to remove a key from a frequency node
func (this *AllOne) removeKeyFromNode(key string, freq int) {
    node := this.freqMap[freq].Value.(*Node)
    delete(node.keys, key)

    // If node is empty, remove it
    if len(node.keys) == 0 {
        this.freqList.Remove(this.freqMap[freq])
        delete(this.freqMap, freq)
    }
}

// Helper function to add a key to a frequency node
func (this *AllOne) addKeyToNode(key string, freq int) {
    var node *list.Element
    if elem, exists := this.freqMap[freq]; exists {
        node = elem
    } else {
        // Create new frequency node
        node = &list.Element{Value: &Node{freq: freq, keys: make(map[string]struct{})}}
        if freq == 1 {
            // Insert at the front if this is the first frequency
            node = this.freqList.PushFront(node)
        } else if prev, exists := this.freqMap[freq-1]; exists {
            // Insert after the previous frequency node
            node = this.freqList.InsertAfter(node, prev)
        }
        this.freqMap[freq] = node
    }
    node.Value.(*Node).keys[key] = struct{}{}
}


/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */