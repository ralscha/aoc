from pathlib import Path

if __name__ == '__main__':
    import igraph
    p = Path("input.txt")
    text = p.read_text(encoding='utf-8')
    lines = text.split('\n')

    verts = set()
    edges = set()

    for line in lines:
        a, b = line.split(': ')
        bs = b.split()

        verts.add(a)
        for b in bs:
            verts.add(b)
            edges.add((a, b))

    g = igraph.Graph()

    for vert in verts:
        g.add_vertex(vert)

    for a, b in edges:
        g.add_edge(a, b)

    cut = g.mincut()

    print(len(cut.partition[0]) * len(cut.partition[1]))
