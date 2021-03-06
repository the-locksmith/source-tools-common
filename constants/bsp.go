package constants

// Common limits
// leaffaces, leafbrushes, planes, and verts are still bounded by
// 16 bit short limits
const MAX_MAP_MODELS = 1024
const MAX_MAP_BRUSHES = 8192
const MAX_MAP_ENTITIES = 8192
const MAX_MAP_TEXINFO = 12288
const MAX_MAP_TEXDATA = 2048
const MAX_MAP_DISPINFO = 2048
const MAX_MAP_DISP_VERTS = MAX_MAP_DISPINFO * ((1 << MAX_MAP_DISP_POWER) + 1) * ((1 << MAX_MAP_DISP_POWER) + 1)
const MAX_MAP_DISP_TRIS = (1 << MAX_MAP_DISP_POWER) * (1 << MAX_MAP_DISP_POWER) * 2
const MAX_DISPVERTS = ((1 << (MAX_MAP_DISP_POWER)) + 1) * ((1 << (MAX_MAP_DISP_POWER)) + 1)
const MAX_DISPTRIS = (1 << (MAX_MAP_DISP_POWER)) * ((1 << (MAX_MAP_DISP_POWER)) * 2)
const MAX_MAP_AREAS = 256
const MAX_MAP_AREA_BYTES = (MAX_MAP_AREAS / 8)
const MAX_MAP_AREAPORTALS = 1024

// Planes come in pairs, thus an even number.
const MAX_MAP_PLANES = 65536
const MAX_MAP_NODES = 65536
const MAX_MAP_BRUSHSIDES = 65536
const MAX_MAP_LEAFS = 65536
const MAX_MAP_VERTS = 65536
const MAX_MAP_VERTNORMALS = 256000
const MAX_MAP_VERTNORMALINDICES = 256000
const MAX_MAP_FACES = 65536
const MAX_MAP_LEAFFACES = 65536
const MAX_MAP_LEAFBRUSHES = 65536
const MAX_MAP_PORTALS = 65536
const MAX_MAP_CLUSTERS = 65536
const MAX_MAP_LEAFWATERDATA = 32768
const MAX_MAP_PORTALVERTS = 128000
const MAX_MAP_EDGES = 256000
const MAX_MAP_SURFEDGES = 512000
const MAX_MAP_LIGHTING = 0x1000000
const MAX_MAP_VISIBILITY = 0x1000000 // increased BSPVERSION 7
const MAX_MAP_TEXTURES = 1024
const MAX_MAP_WORLDLIGHTS = 8192
const MAX_MAP_CUBEMAPSAMPLES = 1024
const MAX_MAP_OVERLAYS = 512
const MAX_MAP_WATEROVERLAYS = 16384
const MAX_MAP_TEXDATA_STRING_DATA = 256000
const MAX_MAP_TEXDATA_STRING_TABLE = 65536
