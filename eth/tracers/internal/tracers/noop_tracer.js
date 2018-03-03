// This file is part of the go-sberex library. The go-sberex library is 
// free software: you can redistribute it and/or modify it under the terms 
// of the GNU Lesser General Public License as published by the Free 
// Software Foundation, either version 3 of the License, or (at your option)
// any later version.
//
// The go-sberex library is distributed in the hope that it will be useful, 
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser 
// General Public License <http://www.gnu.org/licenses/> for more details.

// noopTracer is just the barebone boilerplate code required from a JavaScript
// object to be usable as a transaction tracer.
{
	// step is invoked for every opcode that the VM executes.
	step: function(log, db) { },

	// fault is invoked when the actual execution of an opcode fails.
	fault: function(log, db) { },

	// result is invoked when all the opcodes have been iterated over and returns
	// the final result of the tracing.
	result: function(ctx, db) { }
}
