module solution

pub enum Capacity {
	small
	medium
	large
	extra_large
}

pub struct CapacityPackage {
pub:
	cpu   u32
	memory u32
	size u32
}

type Packages = map[Capacity]CapacityPackage

pub fn get_capacity(cap string) !Capacity {
	match cap {
		'small' {
			return Capacity.small
		}
		'medium' {
			return Capacity.medium
		}
		'large' {
			return Capacity.large
		}
		'extra_large' {
			return Capacity.extra_large
		}
		else {
			return error('invalid capacity')
		}
	}
}